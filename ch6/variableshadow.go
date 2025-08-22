package main

import (
	"errors"
	"fmt"
)

var a int = 2023

func checkYear() error {
	err := errors.New("wrong year")

	switch a, err := getYear(); a {
	case 2023:
		fmt.Println("it is", a, err)
	case 2024:
		fmt.Println("it is", a)
		err = nil
	}
	fmt.Println("after check, it is", a)
	return err
}

type new int

func getYear() (new, error) {
	var b int16 = 2024
	return new(b), nil
}

func main() {
	err := checkYear()
	if err != nil {
		fmt.Println("call checkYear error:", err)
		return
	}
	fmt.Println("call checkYear ok")
}
