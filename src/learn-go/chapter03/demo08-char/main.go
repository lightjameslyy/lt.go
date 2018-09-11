package main

// import "fmt"
// import "unsafe"

import (
	"fmt"
)

// 字符串
func main() {

	var c1 byte = 'a'
	var c2 byte = '0'

	fmt.Println(c1, c2) // 输出 97， 48
	fmt.Printf("%c %c\n", c1, c2)

	// 越界
	// var c3 byte = '北'
	// fmt.Printf("%c3", c3)

	var c4 int = 22269 // '国'的utf-8编码
	fmt.Printf("%c\n", c4)
}
