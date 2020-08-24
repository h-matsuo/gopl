package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

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
	lissajous(w, config)
}

type lissajousConfig struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

func newLissajousConfig() *lissajousConfig {
	return &lissajousConfig{
		cycles:  5,
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
	}
}

func parseQuery(r *http.Request) (*lissajousConfig, error) {
	c := newLissajousConfig()
	// NOTE: Omit strict query validation
	cycles, err := getIntQuery(r, "cycles")
	if cycles > 0 {
		c.cycles = cycles
	} else if err != nil {
		return nil, err
	}
	res, err := getFloatQuery(r, "res")
	if !math.IsNaN(res) {
		c.res = res
	} else if err != nil {
		return nil, err
	}
	size, err := getIntQuery(r, "size")
	if size > 0 {
		c.size = size
	} else if err != nil {
		return nil, err
	}
	nframes, err := getIntQuery(r, "nframes")
	if nframes > 0 {
		c.nframes = nframes
	} else if err != nil {
		return nil, err
	}
	delay, err := getIntQuery(r, "delay")
	if delay > 0 {
		c.delay = delay
	} else if err != nil {
		return nil, err
	}
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

var palette = color.Palette{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer, config *lissajousConfig) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: config.nframes}
	phase := 0.0
	for i := 0; i < config.nframes; i++ {
		rect := image.Rect(0, 0, 2*config.size+1, 2*config.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(config.cycles)*2*math.Pi; t += config.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(config.size+int(x*float64(config.size)+0.5), config.size+int(y*float64(config.size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, config.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
