package main

import "fmt"

func maxString(sl []string) string {
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

func main() {
	fmt.Println(maxString([]string{"11", "22", "44", "66", "77", "10"})) // 输出：77
}
