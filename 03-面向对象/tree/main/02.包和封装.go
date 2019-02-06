package main

import (
	"demo/tree"	//使用导入的包，公开public的方法和变量首字母都要大写
	"fmt"
)

func main() {
	// 创建结构初始化
	var root tree.Node
	left := tree.Node{Val: 1}
	right := tree.Node{2, nil, nil}
	root.Left = &left
	root.Right = &right
	fmt.Println(root)

	//  创建结构切片初始化
	nodes := []tree.Node{
		{Val: 1},
		{Val: 2},
		{3, nil, nil},
	}
	fmt.Println(nodes)

	// 使用工厂函数替代
	node1 := tree.CreateNode(3)
	node2 := tree.CreateNode(4)
	left.Right = node1
	right.Left = node2
	fmt.Println(*node1, node2)

	// 使用结构方法
	root.InOrder()
	fmt.Println()
	(&root).PrintNodeVal()
	root.PrintNodeVal()
	root.SetNodeVal(999)
	root.PrintNodeVal()
	var pNil *tree.Node
	pNil.SetNodeVal(1)
}
