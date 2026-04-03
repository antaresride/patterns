package main

import (
	"testing"
)

func BenchmarkTraditional(b *testing.B) {
	n := 1000
	for b.Loop() {
		for range n {

		}
	}
}

func BenchmarkRange(b *testing.B) {
	n := 1000
	for range b.N {
		for j := range n {
			_ = j
		}
	}
}
