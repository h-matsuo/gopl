package loader

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type format interface {
	Sniff([]byte) (isMatched bool, kind string)
	List([]byte, io.Writer) error
}

var formats []format

func RegisterFormat(f format) {
	formats = append(formats, f)
}

func Load(in io.Reader, out io.Writer) error {
	buf := new(bytes.Buffer)
	io.Copy(buf, in)
	b := buf.Bytes()
	for _, f := range formats {
		if isMatched, kind := f.Sniff(b); isMatched {
			fmt.Fprintln(os.Stderr, "Input format =", kind)
			if err := f.List(b, out); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("unknown format")
}
