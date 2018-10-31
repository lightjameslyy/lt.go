package main

import "fmt"

func main() {
	Q := []int{0}
	//Q = Q[1:]
	//Q = Q[:len(Q)-1]
	for i := 1; i < 10; i++ {
		Q = append(Q, i)
		fmt.Println(Q)
	}
}
