package main

func main() {
	var x interface{} = 13
	switch x.(type) {
	case nil:
		println("x is nil")
	case int:
		println("the type of x is int")
	case string:
		println("the type of x is string")
	case bool:
		println("the type of x is string")
	default:
		println("don't support the type")
	}
}
