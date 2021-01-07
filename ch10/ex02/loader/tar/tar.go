package tar

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"

	"github.com/h-matsuo/gopl/ch10/ex02/loader"
)

type tarFormat struct{}

func init() {
	loader.RegisterFormat(&tarFormat{})
}

func (f *tarFormat) Sniff(b []byte) (isMatched bool, kind string) {
	magicNumber := []byte{'u', 's', 't', 'a', 'r'}
	numBytesToSkip := int64(257)

	readBytes := make([]byte, len(magicNumber))
	if _, err := bytes.NewReader(b).ReadAt(readBytes, numBytesToSkip); err != nil {
		return false, ""
	}

	if bytes.Compare(readBytes, magicNumber) != 0 {
		return false, ""
	}
	return true, "tar"
}

func (f *tarFormat) List(b []byte, out io.Writer) error {
	tr := tar.NewReader(bytes.NewReader(b))
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return err
		}
		fmt.Fprintf(out, "========================================\n")
		fmt.Fprintf(out, "[Contents of %s]\n", hdr.Name)
		if _, err := io.Copy(out, tr); err != nil {
			return err
		}
	}
	return nil
}
