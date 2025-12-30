package main

func main() {
	a, b := false, true
	if a && b != true {
		println("(a && b) != true")
		return
	}
	println("a && (b != true) == false")
}
