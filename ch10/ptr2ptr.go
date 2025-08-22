package main

func main() {
	var a int = 5
	var p1 *int = &a
	println(*p1) // 5
	var b int = 55
	var p2 *int = &b
	println(*p2) // 55

	var pp **int = &p1
	println(**pp) // 5
	pp = &p2
	println(**pp) // 55
}
