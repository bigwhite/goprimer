package main

import "fmt"

func main() {
	var s = "中国人"

	for i := 0; i < len(s); i++ {
		fmt.Printf("index: %d, value: 0x%x\n", i, s[i])
	}

	for i, v := range s {
		fmt.Printf("index: %d, value: 0x%x\n", i, v)
	}
}
