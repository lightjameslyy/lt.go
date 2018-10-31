package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	// go 语言可以返回局部变量的地址
	return &Node{Value: value}
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node, ignored.")
	} else {
		node.Value = value
	}
}
