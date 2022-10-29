package main_test

import (
	"bytes"
	"strings"
	"testing"
)

var result string

func BenchmarkExample(b *testing.B) {
	b.Run(
		"strings.Builder", func(b *testing.B) {
			buf := strings.Builder{}
			buf.Write([]byte("Hello world!"))
			for i := 0; i < b.N; i++ {
				result = buf.String()
			}
		},
	)
	b.Run(
		"bytes.Buffer", func(b *testing.B) {
			buf := bytes.Buffer{}
			buf.Write([]byte("Hello world!"))
			for i := 0; i < b.N; i++ {
				result = buf.String()
			}
		},
	)
}
