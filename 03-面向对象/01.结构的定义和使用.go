package main

import "fmt"

// 定义结构
type treeNode struct {
	val         int
	left, right *treeNode
}

// 定义结构的方法，(node treeNode)表示this，定义和printNode(node treeNode)一致，treeNode值或指针都可以进行调用该方法(自动识别)
func (node treeNode) printNodeVal() {
	fmt.Println(node.val)
}

// go函数参数为传值，需要改值要传引用即指针，treeNode值或指针都可以进行调用该方法(自动识别)
// 指针接收：在改变内容、结构过大、一致性时考虑使用
func (node *treeNode) setNodeVal(val int) {
	if (node == nil) {
		fmt.Println("nil")
		return
	}
	node.val = val
}

// 中序遍历方法
func (node *treeNode) inOrder() {
	if (node == nil) {
		return
	}
	node.left.inOrder()
	fmt.Print(node.val, " ")
	node.right.inOrder()
}

// 没有结构构造函数可用工厂模式替代初始化(返回局部变量指针不会被回收，指针提示编译器自动分配到栈上)
func createTreeNode(val int) *treeNode {
	return &treeNode{val: val}
}

func main() {
	// 创建结构初始化
	var root treeNode
	left := treeNode{val: 1}
	right := treeNode{2, nil, nil}
	root.left = &left
	root.right = &right
	fmt.Println(root)

	//  创建结构切片初始化
	nodes := []treeNode{
		{val: 1},
		{val: 2},
		{3, nil, nil},
	}
	fmt.Println(nodes)

	// 使用工厂函数替代
	node1 := createTreeNode(3)
	node2 := createTreeNode(4)
	left.right = node1
	right.left = node2
	fmt.Println(*node1, node2)

	// 使用结构方法
	root.inOrder()
	fmt.Println()
	(&root).printNodeVal()
	root.printNodeVal()
	root.setNodeVal(999)
	root.printNodeVal()
	var pNil *treeNode
	pNil.setNodeVal(1)
}
