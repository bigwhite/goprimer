package main

import "fmt"

func myAppend(sl []int, elems ...int) []int {
	fmt.Printf("%T\n", elems) // []int
	if len(elems) == 0 {
		println("no elems to append")
		return sl
	}

	sl = append(sl, elems...)
	return sl
}

func main() {
	sl := []int{1, 2, 3}
	sl = myAppend(sl) // no elems to append
	fmt.Println(sl)   // [1 2 3]
	sl = myAppend(sl, 4, 5, 6)
	fmt.Println(sl) // [1 2 3 4 5 6]
}
