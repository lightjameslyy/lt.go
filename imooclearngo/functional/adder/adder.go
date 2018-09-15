package main

import (
	"fmt"
)

// 闭包
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// 正统的函数式编程（只有常量和函数）
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}

	a2 := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a2 = a2(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}
}
