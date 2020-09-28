package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	width, height = 1024, 1024
)

type config struct {
	x          float64
	y          float64
	scale      float64
	iterations int
	xmin       float64
	xmax       float64
	ymin       float64
	ymax       float64
}

func newConfig() *config {
	c := &config{
		x:          0.0,
		y:          0.0,
		scale:      1.0,
		iterations: 100,
	}
	c.xmin, c.xmax = c.x-2.0/c.scale, c.x+2.0/c.scale
	c.ymin, c.ymax = c.y-2.0/c.scale, c.y+2.0/c.scale
	return c
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	c, err := parseQuery(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}
	generateMandelbrot(w, c)
}

func parseQuery(r *http.Request) (*config, error) {
	c := newConfig()
	// NOTE: Omit strict query validation
	x, err := getFloatQuery(r, "x")
	if !math.IsNaN(x) {
		c.x = x
	} else if err != nil {
		return nil, err
	}
	y, err := getFloatQuery(r, "y")
	if !math.IsNaN(y) {
		c.y = y
	} else if err != nil {
		return nil, err
	}
	scale, err := getFloatQuery(r, "scale")
	if !math.IsNaN(scale) {
		c.scale = scale
	} else if err != nil {
		return nil, err
	}
	iterations, err := getIntQuery(r, "iterations")
	if iterations > 0 {
		c.iterations = iterations
	} else if err != nil {
		return nil, err
	}
	c.xmin, c.xmax = c.x-2.0/c.scale, c.x+2.0/c.scale
	c.ymin, c.ymax = c.y-2.0/c.scale, c.y+2.0/c.scale
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

func generateMandelbrot(out io.Writer, c *config) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	yStride := 1.0 / height * (c.ymax - c.ymin)
	xStride := 1.0 / width * (c.xmax - c.xmin)
	for py := 0; py < height; py++ {
		y := float64(py)*yStride + c.ymin
		for px := 0; px < width; px++ {
			x := float64(px)*xStride + c.xmin
			subColor11 := mandelbrot(complex(x+xStride/4, y-yStride/4), c)
			subColor12 := mandelbrot(complex(x+xStride/4, y+yStride/4), c)
			subColor21 := mandelbrot(complex(x-xStride/4, y-yStride/4), c)
			subColor22 := mandelbrot(complex(x-xStride/4, y+yStride/4), c)
			img.Set(px, py, getAverageColor([]color.Color{subColor11, subColor12, subColor21, subColor22}))
		}
	}
	png.Encode(out, img)
}

func mandelbrot(z complex128, c *config) color.Color {
	m := logarithmic(float64(c.iterations))
	var v complex128
	for n := 0; n < c.iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			hue := uint(logarithmic(float64(n)) / m * 360)
			c := HSV{hue, 100, 100}
			return c
		}
	}
	return color.Black
}

func logarithmic(x float64) (y float64) {
	return math.Log10(x + 1)
}
