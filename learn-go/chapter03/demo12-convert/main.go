package main

import (
	"fmt"
	"strconv"
	_ "unsafe" // '-': 不使用该库
)

// basic types conversions
func main() {

	var i int32 = 100
	// i -> float
	// var f float32 = i	// wrong
	var f float32 = float32(i)
	var i8 int8 = int8(i)
	var i64 int64 = int64(i)

	fmt.Printf("i: %v, f: %v, i8: %v, i64: %v\n", i, f, i8, i64)

	// overflow
	var n1 int64 = 99999
	var n2 int8 = int8(n1)
	fmt.Println(n2)

	// 基本数据类型 -> string
	var l1 int = 99
	var l2 float64 = 23.456
	var l3 bool = true
	var l4 byte = 'h'
	var str string

	// 1. fmt.Sprintf()
	str = fmt.Sprintf("%d", l1)
	fmt.Printf("str type: %T, str = %q\n", str, str)
	str = fmt.Sprintf("%f", l2)
	fmt.Printf("str type: %T, str = %q\n", str, str)
	str = fmt.Sprintf("%t", l3)
	fmt.Printf("str type: %T, str = %q\n", str, str)
	str = fmt.Sprintf("%c", l4)
	fmt.Printf("str type: %T, str = %q\n", str, str)

	// 2. strconv
	var m1 int = 99
	var m2 float64 = 23.456
	var m3 bool = true

	str = strconv.FormatInt(int64(m1), 10)
	fmt.Printf("str type: %T, str = %q\n", str, str)
	str = strconv.FormatFloat(m2, 'f', 10, 64)
	fmt.Printf("str type: %T, str = %q\n", str, str)
	str = strconv.FormatBool(m3)
	fmt.Printf("str type: %T, str = %q\n", str, str)

}
