package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("error1")
	err2 := errors.New("error2")
	err3 := errors.New("error3")
	err := errors.Join(err1, err2, err3)
	fmt.Println(err)
}
