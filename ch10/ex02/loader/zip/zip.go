package zip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"

	"github.com/h-matsuo/gopl/ch10/ex02/loader"
)

type zipFormat struct{}

func init() {
	loader.RegisterFormat(&zipFormat{})
}

func (f *zipFormat) Sniff(b []byte) (isMatched bool, kind string) {
	magicNumber := []byte{'P', 'K'}
	numBytesToSkip := int64(0)

	readBytes := make([]byte, len(magicNumber))
	if _, err := bytes.NewReader(b).ReadAt(readBytes, numBytesToSkip); err != nil {
		return false, ""
	}

	if bytes.Compare(readBytes, magicNumber) != 0 {
		return false, ""
	}
	return true, "zip"
}

func (f *zipFormat) List(b []byte, out io.Writer) error {
	zr, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return err
	}
	for _, f := range zr.File {
		fmt.Fprintf(out, "========================================\n")
		fmt.Fprintf(out, "[Contents of %s]\n", f.Name)
		rc, err := f.Open()
		defer rc.Close()
		if err != nil {
			return err
		}
		_, err = io.CopyN(out, rc, 68)
		if err != nil {
			return err
		}
		fmt.Fprintf(out, "\n")
	}
	return nil
}
