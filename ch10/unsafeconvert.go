package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 0x12345678
	var pa *int = &a
	var pb *byte = (*byte)(unsafe.Pointer(pa)) // ok
	fmt.Printf("%x\n", *pb)                    // 78
}
