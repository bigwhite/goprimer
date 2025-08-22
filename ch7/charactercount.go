package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s = "中国人"
	fmt.Println("the character count in s is", utf8.RuneCountInString(s)) // 3

	for _, c := range s {
		fmt.Printf("0x%x ", c) // 0x4e2d 0x56fd 0x4eba
	}
	fmt.Printf("\n")
}
