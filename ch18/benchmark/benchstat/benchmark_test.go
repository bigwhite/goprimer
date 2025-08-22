package main

import (
	"testing"
)

func BenchmarkStrcat(b *testing.B) {
	str1 := "Hello"
	str2 := "World"
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := Strcat(str1, str2)
		_ = result
	}
}
