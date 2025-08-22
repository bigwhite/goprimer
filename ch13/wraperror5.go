package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("error1")
	err2 := fmt.Errorf("error2: wrap %w", err1)
	err3 := fmt.Errorf("error3: wrap %w", err2)
	err := fmt.Errorf("error: wrap %w", err3)
	fmt.Println(errors.Is(err, err1))  // true
	fmt.Println(errors.Is(err, err2))  // true
	fmt.Println(errors.Is(err, err3))  // true
	fmt.Println(errors.Is(err2, err1)) // true
	fmt.Println(errors.Is(err3, err1)) // true
	fmt.Println(errors.Is(err3, err2)) // true
}
