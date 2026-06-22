package main

import "testing"

func BenchmarkParallel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parallel()
	}
}

func BenchmarkSequential(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sequential()
	}
}
