package main

import (
	"fmt"
)

func main() {
	mp := make(map[int]int)
	mp[1] = 3
	mp[40] = 9
	mp[2] = 23
	mp[0] = 344

	for k, v := range mp {
		fmt.Println(k, v)
	}
}
