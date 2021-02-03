package bzip

import (
	"bytes"
	"io"
	"log"
	"os/exec"
)

type writer struct {
	w io.Writer // underlying output stream
}

func NewWriter(out io.Writer) io.WriteCloser {
	w := &writer{w: out}
	return w
}

func (w *writer) Write(data []byte) (int, error) {
	cmd := exec.Command("bzip2")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdin.Close()
		io.Copy(stdin, bytes.NewReader(data))
	}()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdout.Close()
		io.Copy(w.w, stdout)
	}()

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return len(data), nil
}

func (w *writer) Close() error {
	return nil
}
