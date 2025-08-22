package main

type T struct {
	n int
	s string
}

func main() {
	var t = T{
		n: 17,
		s: "hello, interface",
	}

	var ei interface{} = t // Go运行时使用eface结构表示ei
	_ = ei
}
