package main

import (
	"fmt"
)

type call struct {
	val interface{}
	err error
}

func main() {
	m := make(map[string]*call)
	c := new(call)

	m["key"] = c
	fmt.Println("m:", m)
	fmt.Println("c:", c)

	// delete函数执行之后，c的空间并没有被释放
	// 由此可见，delete是“按值删除”
	delete(m, "key")
	fmt.Println("m:", m)
	fmt.Println("c:", c)
}
