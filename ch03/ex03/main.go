package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	// Calculate all coordinates
	points := [][9]float64{}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, ax, ay := corner(i+1, j)
			z, bx, by := corner(i, j)
			_, cx, cy := corner(i, j+1)
			_, dx, dy := corner(i+1, j+1)
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

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for _, point := range points {
		ax, ay := point[1], point[2]
		bx, by := point[3], point[4]
		cx, cy := point[5], point[6]
		dx, dy := point[7], point[8]
		colorCode := computeHSLColor(point[0], max, min)
		fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s' />\n",
			ax, ay, bx, by, cx, cy, dx, dy, colorCode)
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (z, sx, sy float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z = f(x, y)

	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
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
