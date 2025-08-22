package main

import (
	"errors"
	"fmt"
)

func rootCause(err error) error {
	for {
		e, ok := err.(interface{ Unwrap() error })
		if !ok {
			return err
		}
		err = e.Unwrap()
		if err == nil {
			return nil
		}
	}
}

func main() {
	err1 := errors.New("error1")
	err2 := fmt.Errorf("error2: wrap %w", err1)
	err3 := fmt.Errorf("error3: wrap %w", err2)
	err := fmt.Errorf("error: wrap %w", err3)
	fmt.Println(err)
	fmt.Println("root cause is", rootCause(err))
	fmt.Println(rootCause(err) == err1)
}
