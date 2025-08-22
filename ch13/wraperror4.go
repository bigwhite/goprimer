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
	fmt.Println(err == err1)
}
