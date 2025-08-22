package benchmark

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkDemo(b *testing.B) {
	fmt.Println("b.N=", b.N)
	for i := 0; i < b.N; i++ {
		rand.Int31()
	}
}
