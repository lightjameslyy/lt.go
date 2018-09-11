package main

import (
	"fmt"
	"unsafe"
)

// bool
func main() {

	b := false
	fmt.Println("b=", b)
	// 1. bool size is 1 byte
	fmt.Printf("size of b: %d\n", unsafe.Sizeof(b))
	// 2. bool赋值时只能赋true或false
	// b = 1
}
