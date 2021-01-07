package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var (
	out = flag.String("out", "jpeg", "outout format (gif, png, jpeg)")
)

func main() {
	flag.Parse()
	var err error
	switch *out {
	case "gif":
		err = toGIF(decodeImage(os.Stdin), os.Stdout)
	case "png":
		err = toPNG(decodeImage(os.Stdin), os.Stdout)
	case "jpeg":
		err = toJPEG(decodeImage(os.Stdin), os.Stdout)
	default:
		fmt.Fprintf(os.Stderr, "unknown output format: %s\n", *out)
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func decodeImage(in io.Reader) image.Image {
	img, kind, err := image.Decode(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return img
}

func toGIF(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, &gif.Options{})
}

func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
