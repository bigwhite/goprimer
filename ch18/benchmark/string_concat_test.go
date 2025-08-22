package benchmark

import (
	"fmt"
	"strings"
	"testing"
)

// 方法一：使用 + 运算符进行连接
func BenchmarkStringConcatenation(b *testing.B) {
	str1 := "Hello"
	str2 := "World"
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := str1 + str2
		_ = result
	}
}

// 方法二：使用 strings.Join 进行连接
func BenchmarkStringsJoin(b *testing.B) {
	str1 := "Hello"
	str2 := "World"
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := strings.Join([]string{str1, str2}, "")
		_ = result
	}
}

// 方法三：使用 fmt.Sprintf 进行连接
func BenchmarkFmtSprintf(b *testing.B) {
	str1 := "Hello"
	str2 := "World"
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := fmt.Sprintf("%s%s", str1, str2)
		_ = result
	}
}

// 方法四：使用 strings.Builder 进行连接
func BenchmarkStringsBuilder(b *testing.B) {
	str1 := "Hello"
	str2 := "World"
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.WriteString(str1)
		builder.WriteString(str2)
		result := builder.String()
		_ = result
	}
}
