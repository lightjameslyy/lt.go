package main

import "fmt"

func main() {
	// golang变量使用方式

	// 第一种：指定变量类型，声明后不赋值，使用默认值
	// int的默认值是0
	var i int
	fmt.Println("i =", i)

	// 第二种：类型推导
	var num = 10.11
	fmt.Println("num =", num)

	// 第三种：使用:=
	// 注意：:=左侧的变量必须是未声明过的，否则会编译出错
	// 下面的方式等价于：
	// var name string
	// name = "tom"
	name := "tom"
	fmt.Println("name =", name)
}
