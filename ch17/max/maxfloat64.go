package main

import "fmt"

func maxFloat(sl []float64) float64 {
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
	fmt.Println(maxFloat([]float64{1.01, 2.02, 3.03, 5.05, 7.07, 0.01})) // 输出：7.07
}
