package main

import "fmt"

func main() {
	var a = "initial" // type inference
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // default value: zero-valued
	fmt.Println(e)

	f := "short" // same as var f string = "short"
	fmt.Println(f)
}
