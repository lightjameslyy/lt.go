package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 普通调用
	f("direct")

	// 开一个goroutine，调用f()
	go f("goroutine")

	// goroutine调用匿名函数
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 等待goroutine运行完
	fmt.Scanln()
	fmt.Println("done")

}
