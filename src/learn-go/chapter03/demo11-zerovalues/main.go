package main

import (
	"fmt"
)

// zero values
func main() {

	var a int     // 0
	var b float32 // 0
	var c float64 // 0
	var d bool    // false
	var e string  // ""
	// %v: 按照变量的值输出
	fmt.Printf("a: %d, b: %v, c: %v, d: %v, e: %v\n", a, b, c, d, e)
}
