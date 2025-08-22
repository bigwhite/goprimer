package main

import (
	"fmt"
	"os"
)

func readFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func processFile(filename string) error {
	data, err := readFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %s", filename)
	}
	fmt.Println(string(data))
	return nil
}

func main() {
	err := processFile("demo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
}
