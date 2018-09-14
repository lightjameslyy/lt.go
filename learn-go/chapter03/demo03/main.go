package main

import "fmt"

// global variable
var g1 = 100
var g2 = "g2"

// 上面也可以写成：
var (
	g3 = 300
	g4 = "g4"
)

func main() {
	// 一次性声明多个变量
	// 第一种声明方式
	var n1, n2, n3 int
	fmt.Println("n1:", n1, ",n2:", n2, ",n3:", n3)

	// 第二种声明方式
	var m1, name, m2 = 100, "james", 888
	fmt.Println("m1:", m1, "name:", name, "m2:", m2)

	// 第三种声明方式
	l1, tag, l2 := 666, "light", 999
	fmt.Println("l1:", l1, "tag:", tag, "l2:", l2)

	// 全局变量
	fmt.Println("g1:", g1, "g2:", g2)
	fmt.Println("g3:", g3, "g4:", g4)
}
