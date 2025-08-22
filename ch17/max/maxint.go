package main

import "fmt"

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

func main() {
	fmt.Println(maxInt([]int{1, 2, -4, -6, 7, 0})) // 输出：7
}
