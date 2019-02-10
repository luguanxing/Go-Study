package main

import (
	"demo/tree" //使用导入的包，公开public的方法和变量首字母都要大写
	"fmt"
)

func main() {
	// 创建结构初始化
	var root tree.Node
	left := tree.Node{Val: 1}
	right := tree.Node{2, nil, nil}
	root.Left = &left
	root.Right = &right

	// 使用工厂函数替代
	node1 := tree.CreateNode(3)
	node2 := tree.CreateNode(4)
	left.Right = node1
	right.Left = node2

	// 测试先序遍历
	root.InOrder()
	fmt.Println()
	// 测试在遍历中使用自己的逻辑函数
	root.InOrderFunc(func(node *tree.Node) {
		node.PrintNodeVal()
		node.Val = 1 // 指针可改值
	})
	root.InOrderFunc(func(node *tree.Node) {
		fmt.Print(node.Val, " ")
	})
	fmt.Println()

	nodeCount := 0
	root.InOrderFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("共有", nodeCount, "个节点")
}
