package main

import (
	"image/color"
	"math"
)

type HSV struct {
	Hue, Saturation, Value uint
}

func (h HSV) RGBA() (r, g, b, a uint32) {
	if h.Hue >= 360 {
		panic("0° <= Hue < 360°")
	}
	if h.Saturation > 100 {
		panic("0%% <= Saturation <= 100%%")
	}
	if h.Value > 100 {
		panic("0%% <= Value <= 100%%")
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
