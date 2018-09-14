package main

// import "fmt"
// import "unsafe"

import (
	"fmt"
)

// 浮点类型
func main() {

	var price float32 = 89.12
	fmt.Println("price =", price)

	var n1 float32 = 0.000043
	var n2 float64 = -0.83343429
	fmt.Println("n1:", n1, "n2:", n2)

	// 精度损失
	var n3 float32 = 123.0000091
	var n4 float64 = -123.0000091
	fmt.Println("n3 =", n3, "n4 =", n4)

	// Golang中浮点类型默认是float64
	n5 := 0.2354
	fmt.Printf("default float type in Golang: %T\n", n5)

}
