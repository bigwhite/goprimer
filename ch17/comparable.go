package main

type foo struct {
	a int
	s string
}

type bar struct {
	a  int
	sl []string
}

func doSomething[T comparable](t T) T {
	var a T
	if a == t {
	}

	if a != t {
	}
	return a
}

func main() {
	doSomething(true)
	doSomething(3)
	doSomething(3.14)
	doSomething(3 + 4i)
	doSomething("hello")
	var p *int
	doSomething(p)
	doSomething(make(chan int))
	doSomething([3]int{1, 2, 3})
	doSomething(foo{})
	//doSomething(bar{}) //  bar does not implement comparable
}
