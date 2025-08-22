package main

import "unsafe"

func main() {
	var arr = [5]int{11, 12, 13, 14, 15}
	var p *int = &arr[0]
	var i uintptr
	for i = 0; i < uintptr(len(arr)); i++ {
		p1 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i*unsafe.Sizeof(*p)))
		println(*p1)
	}
}
