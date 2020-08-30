package echo

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const lenOsArgs = 100

var osArgsBuf []string

func init() {
	osArgsBuf = make([]string, 0, lenOsArgs)
	osArgsBuf = append(osArgsBuf, "cmd-name")
	for i := 1; i < lenOsArgs; i++ {
		osArgsBuf = append(osArgsBuf, fmt.Sprintf("arg[%d]", i))
	}
}

func BenchmarkNonEfficientEcho(b *testing.B) {
	out = ioutil.Discard
	osArgs = osArgsBuf
	for i := 0; i < b.N; i++ {
		nonEfficientEcho()
	}
}

func BenchmarkEfficientEcho(b *testing.B) {
	out = ioutil.Discard
	osArgs = osArgsBuf
	for i := 0; i < b.N; i++ {
		efficientEcho()
	}
}
