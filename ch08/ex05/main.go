package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	yStride := 1.0 / height * (ymax - ymin)
	xStride := 1.0 / width * (xmax - xmin)
	var wg sync.WaitGroup
	for py := 0; py < height; py++ {
		y := float64(py)*yStride + ymin
		for px := 0; px < width; px++ {
			x := float64(px)*xStride + xmin
			wg.Add(1)
			go func(px, py int) {
				defer wg.Done()
				subColor11 := mandelbrot(complex(x+xStride/4, y-yStride/4))
				subColor12 := mandelbrot(complex(x+xStride/4, y+yStride/4))
				subColor21 := mandelbrot(complex(x-xStride/4, y-yStride/4))
				subColor22 := mandelbrot(complex(x-xStride/4, y+yStride/4))
				img.Set(px, py, getAverageColor([]color.Color{subColor11, subColor12, subColor21, subColor22}))
			}(px, py)
		}
	}
	wg.Wait()
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 50

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			hue := uint(float64(n) / float64(iterations) * 360)
			c := HSV{hue, 100, 100}
			return c
		}
	}
	return color.Black
}
