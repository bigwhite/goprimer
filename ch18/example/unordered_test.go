package main

import "fmt"

func ExampleUnordered() {
	m := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}
	for _, v := range m {
		fmt.Println(v)
	}
	// Unordered output: 4
	// 2
	// 1
	// 3
	// 0
}
