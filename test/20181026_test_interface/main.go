package main

import "fmt"

func main() {
	var c cat
	//值作为参数传递
	invoke(&c)
}

//需要一个animal接口作为参数
func invoke(a animal) {
	a.printInfo()
}

type animal interface {
	printInfo()
}

type cat int

//指针接收者实现animal接口
func (c *cat) printInfo() {
	fmt.Println("a cat")
}
