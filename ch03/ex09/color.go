package main

import (
	"fmt"
	"image/color"
	"math"
)

type HSV struct {
	Hue, Saturation, Value uint
}

func (h HSV) RGBA() (r, g, b, a uint32) {
	if h.Hue >= 360 {
		panic(fmt.Sprintf("0° <= Hue < 360°, actual Hue: %d", h.Hue))
	}
	if h.Saturation > 100 {
		panic(fmt.Sprintf("0%% <= Saturation <= 100%%, actual Saturation: %d", h.Saturation))
	}
	if h.Value > 100 {
		panic(fmt.Sprintf("0%% <= Value <= 100%%, actual Value: %d", h.Value))
	}
	s := float64(h.Saturation) / 100
	v := float64(h.Value) / 100

	c := v * s
	x := c * (1 - math.Abs(math.Mod(float64(h.Hue)/60, 2)-1))
	m := v - c
	var r2, g2, b2 float64
	switch {
	case h.Hue < 60:
		r2, g2, b2 = c, x, 0
	case h.Hue < 120:
		r2, g2, b2 = x, c, 0
	case h.Hue < 180:
		r2, g2, b2 = 0, c, x
	case h.Hue < 240:
		r2, g2, b2 = 0, x, c
	case h.Hue < 300:
		r2, g2, b2 = x, 0, c
	case h.Hue < 360:
		r2, g2, b2 = c, 0, x
	}
	return color.RGBA{uint8((r2 + m) * 255), uint8((g2 + m) * 255), uint8((b2 + m) * 255), 0xff}.RGBA()
}

func getAverageColor(colors []color.Color) color.Color {
	r, g, b, a := uint64(0), uint64(0), uint64(0), uint64(0)
	for _, c := range colors {
		r2, g2, b2, a2 := c.RGBA()
		r += uint64(r2)
		g += uint64(g2)
		b += uint64(b2)
		a += uint64(a2)
	}
	l := uint64(len(colors))
	return tempColor{uint32(r / l), uint32(g / l), uint32(b / l), uint32(a / l)}
}

type tempColor struct {
	r, g, b, a uint32
}

func (c tempColor) RGBA() (r, g, b, a uint32) {
	return c.r, c.g, c.b, c.a
}
