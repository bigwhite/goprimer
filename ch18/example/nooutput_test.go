package main

import "fmt"

func ExampleNooutput() {
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
	// Output:
}
