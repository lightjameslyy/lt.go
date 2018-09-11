package main

import (
	"fmt"
)

// string
func main() {

	var address string = "北京海淀"
	fmt.Println(address)

	// 字符串一旦赋值，不能修改
	// var str = "hello"
	// cannot assign to str[0]
	// str[0] = 'a'

	// 双引号和反引号
	str1 := "abc\nabc"
	fmt.Println("双引号：\n", str1)

	str2 := `
	package main

	import (
		"fmt"
	)
	
	// string
	func main() {
	
		var address string = "北京海淀"
		fmt.Println(address)
	
		// 字符串一旦赋值，不能修改
		// var str = "hello"
		// cannot assign to str[0]
		// str[0] = 'a'
	
		// 双引号和反引号
		str1 := "abc\nabc"
		fmt.Println(str1)
	}	
	`
	fmt.Println("反引号：\n", str2)

	// 字符串拼接
	var str3 = "hello" + "world"
	str3 += " haha!"
	fmt.Println(str3)

	// 多行拼接，”+“要放在行末
	var str4 = "a" + "a" + "a" +
		"a" + "a" + "a"
	// 下面的写法是错的
	// var str4 = "a" + "a" + "a"
	// +"a" + "a" + "a"
	fmt.Println(str4)
}
