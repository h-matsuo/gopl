package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

type svgConfig struct {
	width   int
	height  int
	cells   int
	xyrange float64
	xyscale float64
	zscale  float64
	angle   float64
}

func newSVGConfig() *svgConfig {
	c := &svgConfig{
		width:   600,
		height:  320,
		cells:   100,
		xyrange: 30.0,
		angle:   math.Pi / 6,
	}
	c.xyscale = float64(c.width) / float64(2) / c.xyrange
	c.zscale = float64(c.height) * 0.4
	return c
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	config, err := parseQuery(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	generateSVG(w, config)
}

func parseQuery(r *http.Request) (*svgConfig, error) {
	c := newSVGConfig()
	// NOTE: Omit strict query validation
	width, err := getIntQuery(r, "width")
	if width > 0 {
		c.width = width
	} else if err != nil {
		return nil, err
	}
	height, err := getIntQuery(r, "height")
	if height > 0 {
		c.height = height
	} else if err != nil {
		return nil, err
	}
	cells, err := getIntQuery(r, "cells")
	if cells > 0 {
		c.cells = cells
	} else if err != nil {
		return nil, err
	}
	xyrange, err := getFloatQuery(r, "xyrange")
	if !math.IsNaN(xyrange) {
		c.xyrange = xyrange
	} else if err != nil {
		return nil, err
	}
	angle, err := getFloatQuery(r, "angle")
	if !math.IsNaN(angle) {
		c.angle = angle
	} else if err != nil {
		return nil, err
	}
	c.xyscale = float64(c.width) / float64(2) / c.xyrange
	c.zscale = float64(c.height) * 0.4
	return c, nil
}

func getIntQuery(r *http.Request, queryName string) (int, error) {
	if valueStr := r.URL.Query().Get(queryName); valueStr != "" {
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return -1, err
		}
		return value, nil
	}
	return -1, nil
}

func getFloatQuery(r *http.Request, queryName string) (float64, error) {
	if valueStr := r.URL.Query().Get(queryName); valueStr != "" {
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return math.NaN(), err
		}
		return value, nil
	}
	return math.NaN(), nil
}

////

func generateSVG(out io.Writer, config *svgConfig) {
	// Calculate all coordinates
	points := [][9]float64{}
	for i := 0; i < config.cells; i++ {
		for j := 0; j < config.cells; j++ {
			_, ax, ay := corner(i+1, j, config)
			z, bx, by := corner(i, j, config)
			_, cx, cy := corner(i, j+1, config)
			_, dx, dy := corner(i+1, j+1, config)
			if isInvalidCoordinate(ax, ay) || isInvalidCoordinate(bx, by) ||
				isInvalidCoordinate(cx, cy) || isInvalidCoordinate(dx, dy) {
				fmt.Fprintf(os.Stderr, "IGNORED: <polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
				continue
			}
			p := [9]float64{z, ax, ay, bx, by, cx, cy, dx, dy}
			points = append(points, p)
		}
	}

	max, min := findMaxAndMin(points)

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", config.width, config.height)
	for _, point := range points {
		ax, ay := point[1], point[2]
		bx, by := point[3], point[4]
		cx, cy := point[5], point[6]
		dx, dy := point[7], point[8]
		colorCode := computeHSLColor(point[0], max, min)
		fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s' />\n",
			ax, ay, bx, by, cx, cy, dx, dy, colorCode)
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int, config *svgConfig) (z, sx, sy float64) {
	x := config.xyrange * (float64(i)/float64(config.cells) - 0.5)
	y := config.xyrange * (float64(j)/float64(config.cells) - 0.5)

	z = f(x, y)

	sx = float64(config.width)/2.0 + (x-y)*math.Cos(config.angle)*config.xyscale
	sy = float64(config.height)/2.0 + (x+y)*math.Sin(config.angle)*config.xyscale - z*config.zscale
	return z, sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func isInvalidCoordinate(x, y float64) bool {
	return math.IsNaN(x) || math.IsInf(x, 0) || math.IsNaN(y) || math.IsInf(y, 0)
}

func findMaxAndMin(points [][9]float64) (max, min float64) {
	max, min = math.Inf(-1), math.Inf(1)
	for _, point := range points {
		z := point[0]
		if z > max {
			max = z
		}
		if z < min {
			min = z
		}
	}
	return max, min
}

func computeHSLColor(z, max, min float64) string {
	hueAngle := 120
	if z >= 0 {
		p := z / max
		hueAngle = 120 - int(120*p)
	} else {
		p := -z / -min
		hueAngle = 120 + int(120*p)
	}
	fmt.Fprintf(os.Stderr, "hueAngle = %d°, z = %g\n", hueAngle, z)
	return fmt.Sprintf("hsl(%d, 100%%, 50%%)", hueAngle)
}
