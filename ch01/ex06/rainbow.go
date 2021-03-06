package main

import (
	"image/color"
	"math"
)

const maxNumColorsInPalette = math.MaxUint8 // Type of Paletted.Pix is uint8
const numColorRepetition = 5

var numColorsBetweenHueColors = int((maxNumColorsInPalette - 6) / 6)

// nextColorIndex is a high-order function to produce an appropriate index
// pointing the color in the palette generated by generateRainbowPalette()
func nextColorIndex() func() uint8 {
	numColors := numColorsBetweenHueColors*6 + 6
	idx := -1
	numCurrentRepetition := 0
	return func() uint8 {
		numCurrentRepetition++
		if numCurrentRepetition >= numColorRepetition {
			idx++
			numCurrentRepetition = 0
			if idx > numColors {
				idx = 0
			}
		}
		return uint8(idx + 1) // +1 for the background color (Black)
	}
}

func generateRainbowPalette() color.Palette {
	palette := make([]color.Color, 0, maxNumColorsInPalette)
	var r, g, b uint8
	diff := 0xff / (numColorsBetweenHueColors + 2)

	// Red: #ff0000
	r, g, b = 0xff, 0x00, 0x00
	palette = append(palette, color.RGBA{r, g, b, 0xff})

	// Gradient
	for i := 1; i <= numColorsBetweenHueColors; i++ {
		g = uint8(0x00 + diff*i)
		palette = append(palette, color.RGBA{r, g, b, 0xff})
	}

	// Yellow: #ffff00
	r, g, b = 0xff, 0xff, 0x00
	palette = append(palette, color.RGBA{r, g, b, 0xff})

	// Gradient
	for i := 1; i <= numColorsBetweenHueColors; i++ {
		r = uint8(0xff - diff*i)
		palette = append(palette, color.RGBA{r, g, b, 0xff})
	}

	// Green: #00ff00
	r, g, b = 0x00, 0xff, 0x00
	palette = append(palette, color.RGBA{r, g, b, 0xff})

	// Gradient
	for i := 1; i <= numColorsBetweenHueColors; i++ {
		b = uint8(0x00 + diff*i)
		palette = append(palette, color.RGBA{r, g, b, 0xff})
	}

	// Light blue: #00ffff
	r, g, b = 0x00, 0xff, 0xff
	palette = append(palette, color.RGBA{r, g, b, 0xff})

	// Gradient
	for i := 1; i <= numColorsBetweenHueColors; i++ {
		g = uint8(0xff - diff*i)
		palette = append(palette, color.RGBA{r, g, b, 0xff})
	}

	// Blue: #0000ff
	r, g, b = 0x00, 0x00, 0xff
	palette = append(palette, color.RGBA{r, g, b, 0xff})

	// Gradient
	for i := 1; i <= numColorsBetweenHueColors; i++ {
		r = uint8(0x00 + diff*i)
		palette = append(palette, color.RGBA{r, g, b, 0xff})
	}

	// Pink: #ff00ff
	r, g, b = 0xff, 0x00, 0xff
	palette = append(palette, color.RGBA{r, g, b, 0xff})

	// Gradient
	for i := 1; i <= numColorsBetweenHueColors; i++ {
		b = uint8(0xff - diff*i)
		palette = append(palette, color.RGBA{r, g, b, 0xff})
	}

	return palette
}
