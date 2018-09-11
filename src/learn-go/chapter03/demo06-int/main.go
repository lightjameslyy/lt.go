package main

// import "fmt"
// import "unsafe"

import (
	"fmt"
	"unsafe"
)

func main() {

	// int8: -128~127
	// var i int8 = -129
	// 其他int类型以此类推
	var i int8 = -128
	fmt.Println(i)

	// uint8: 0~255
	// var j uint8 = -1
	var j uint8 = 255
	fmt.Println(j)

	// int, uint, rune, byte
	// rune = int32
	// byte = uint8
	var a int = 8900
	var b uint = 1
	var c byte = 96
	fmt.Println(a, b, c)

	// 查看数据类型
	var n1 = 100
	fmt.Printf("type of n1: %T\n", n1)

	// 查看字节数
	var n2 int64 = 10
	fmt.Printf("type of n2: %T, size of n2: %d\n", n2, unsafe.Sizeof(n2))
}
