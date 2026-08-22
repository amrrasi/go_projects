package utils

import "testing"

func BenchmarkConcatStringWithoutBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringWithoutBuilder("test", "salam")
	}
}

func BenchmarkConcatStringWithBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringWithBuilder("test", "salam")
	}
}
