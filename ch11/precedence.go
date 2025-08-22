package main

func main() {
	a, b := true, false
	if a && b != true {
		println("(a && b) != true")
		return
	}
	println("a && (b != true) == true")
}
