package main

/*
	命名返回值 and 不带参数的return
*/

import "fmt"

func Fetch(key interface{}) (interface{}, bool) {
	// 匿名返回值没有初始化
	// return  // error
	return nil, false
}

func Get(key interface{}) (value interface{}, ok bool) {
	// 命名返回值会用默认值初始化
	// 直接返回时，返回值为对应数据类型的默认值
	return
}

func main() {
	value, ok := Get("test")
	fmt.Println(value, ok)

	value, ok = Fetch("test")
	fmt.Println(value, ok)
}
