package main

import "fmt"

func main() {
	str := `This is a raw string with a backquote: "` + "`" + `" inside.`
	fmt.Println(str)
}
