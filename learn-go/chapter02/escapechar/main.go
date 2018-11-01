package main

import "fmt"	//格式化，io

func main() {
	// test escape char

	fmt.Println("tom\tjack")

	fmt.Println("hello\nworld")

	fmt.Println("C:\\Users\\Administrator\\Desktop")

	fmt.Println("tom say: \"i love you\"")

	// \r 回车，从当前行的最前面开始输出，覆盖以前的内容
	fmt.Println("hello world\rbingo")

	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
}