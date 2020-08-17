package echo

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const lenOsArgs = 100

var osArgsBuf = generateOsArgs()

func generateOsArgs() []string {
	result := make([]string, 0, lenOsArgs)
	result = append(result, "cmd-name")
	for i := 1; i < lenOsArgs; i++ {
		result = append(result, fmt.Sprintf("arg[%d]", i))
	}
	return result
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
