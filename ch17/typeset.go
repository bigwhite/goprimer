package main

type Intf1 interface {
	~int | string
	F1()
	F2()
}

type Intf2 interface {
	~int | ~float64
}

type I interface {
	Intf1
	M1()
	M2()
	int | ~string | Intf2
}

func doSomething[T I](t T) {
}

type MyInt int

func (MyInt) F1() {
}
func (MyInt) F2() {
}
func (MyInt) M1() {
}
func (MyInt) M2() {
}

func main() {
	var a int = 11
	doSomething(a) //int does not implement I (missing F1 method)

	var b = MyInt(a)
	doSomething(b) // ok
}
