package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("error1")
	err2 := errors.New("error2")
	err3 := errors.New("error3")
	err := fmt.Errorf("error: wrap %w, %w, %w", err1, err2, err3)
	fmt.Println(errors.Is(err, err1)) // true
	fmt.Println(errors.Is(err, err2)) // true
	fmt.Println(errors.Is(err, err3)) // true

	err = errors.Join(err1, err2, err3)
	fmt.Println(errors.Is(err, err1)) // true
	fmt.Println(errors.Is(err, err2)) // true
	fmt.Println(errors.Is(err, err3)) // true
}
