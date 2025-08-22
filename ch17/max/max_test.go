package main

import "testing"

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func maxGenerics[T ordered](sl []T) T {
	if len(sl) == 0 {
		panic("slice is empty")
	}

	max := sl[0]
	for _, v := range sl[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func maxAny(sl []any) any {
	if len(sl) == 0 {
		panic("slice is empty")
	}

	max := sl[0]
	for _, v := range sl[1:] {
		switch v.(type) {
		case int:
			if v.(int) > max.(int) {
				max = v
			}
		case string:
			if v.(string) > max.(string) {
				max = v
			}
		case float64:
			if v.(float64) > max.(float64) {
				max = v
			}
		}
	}
	return max
}

func maxInt(sl []int) int {
	if len(sl) == 0 {
		panic("slice is empty")
	}

	max := sl[0]
	for _, v := range sl[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func BenchmarkMaxInt(b *testing.B) {
	sl := []int{1, 2, 3, 4, 7, 8, 9, 0}
	for i := 0; i < b.N; i++ {
		maxInt(sl)
	}
}

func BenchmarkMaxAny(b *testing.B) {
	sl := []any{1, 2, 3, 4, 7, 8, 9, 0}
	for i := 0; i < b.N; i++ {
		maxAny(sl)
	}
}

func BenchmarkMaxGenerics(b *testing.B) {
	sl := []int{1, 2, 3, 4, 7, 8, 9, 0}
	for i := 0; i < b.N; i++ {
		maxGenerics(sl)
	}
}
