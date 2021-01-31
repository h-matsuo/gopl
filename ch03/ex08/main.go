package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
	"math/cmplx"
	"os"
)

const (
	// out5
	// x, y          = -0.10812, -0.90445 // Central coordinate
	// scale         = 20000.0
	// width, height = 1024, 1024
	// iterations    = 10000
	// out7
	// x, y          = -0.10811, -0.9046 // Central coordinate
	// scale         = 20000.0
	// width, height = 1024, 1024
	// iterations    = 1000
	// out9
	// x, y          = -0.1081092, -0.9045997 // Central coordinate
	// scale         = 1000000.0
	// width, height = 1024, 1024
	// iterations    = 5000
	// out10
	// x, y          = -0.1081092, -0.904599679869 // Central coordinate
	// scale         = 10000000000.0
	// width, height = 1024, 1024
	// iterations    = 5000
	// out11
	// x, y          = -0.10810919999, -0.904599679869 // Central coordinate
	// scale         = 1000000000000.0
	// width, height = 1024, 1024
	// iterations    = 5000

	// // out12
	// x, y          = -0.10812, -0.90445 // Central coordinate
	// scale         = 50000.0
	// width, height = 1024, 1024
	// iterations    = 10000

	// // out13
	// x, y          = -0.10812, -0.904475 // Central coordinate
	// scale         = 1000000.0
	// width, height = 1024, 1024
	// iterations    = 1000

	// // out14
	// x, y          = -0.10812, -0.904475 // Central coordinate
	// scale         = 1000000.0
	// width, height = 1024, 1024
	// iterations    = 10000

	// out15
	x, y          = -.743030, .126433 // Central coordinate
	scale         = 190.99
	width, height = 1024, 1024
	iterations    = 10000

	xmin, xmax = x - 2.0/scale, x + 2.0/scale
	ymin, ymax = y - 2.0/scale, y + 2.0/scale
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	yStride := 1.0 / height * (ymax - ymin)
	xStride := 1.0 / width * (xmax - xmin)
	for py := 0; py < height; py++ {
		y := float64(py)*yStride + ymin
		for px := 0; px < width; px++ {
			x := float64(px)*xStride + xmin
			subColor11 := mandelbrot(complex(x+xStride/4, y-yStride/4))
			subColor12 := mandelbrot(complex(x+xStride/4, y+yStride/4))
			subColor21 := mandelbrot(complex(x-xStride/4, y-yStride/4))
			subColor22 := mandelbrot(complex(x-xStride/4, y+yStride/4))
			img.Set(px, py, getAverageColor([]color.Color{subColor11, subColor12, subColor21, subColor22}))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(r, i big.Float) color.Color {
	m := logarithmic(float64(iterations))
	var v complex128
	for n := 0; n < iterations; n++ {
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
