package main

import (
	"strings"
	"testing"
)

var s = strings.Repeat("a", 1024)

func test() {
	b := []byte(s)
	_ = string(b)
}


func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

