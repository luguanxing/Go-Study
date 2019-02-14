package tree

import "fmt"

// 定义结构的方法，(node treeNode)表示this，定义和printNode(node treeNode)一致，treeNode值或指针都可以进行调用该方法(自动识别)
func (node Node) PrintNodeVal() {
	fmt.Println(node.Val)
}

// go函数参数为传值，需要改值要传引用即指针，treeNode值或指针都可以进行调用该方法(自动识别)
// 指针接收：在改变内容、结构过大、一致性时考虑使用
func (node *Node) SetNodeVal(val int) {
	if (node == nil) {
		fmt.Println("nil")
		return
	}
	node.Val = val
}

// 中序遍历方法
func (node *Node) InOrder() {
	if (node == nil) {
		return
	}
	node.Left.InOrder()
	fmt.Print(node.Val, " ")
	node.Right.InOrder()
}

// 没有结构构造函数可用工厂模式替代初始化(返回局部变量指针不会被回收，指针提示编译器自动分配到栈上)
func CreateNode(val int) *Node {
	return &Node{Val: val}
}

// 中序遍历传递方法
func (node *Node) InOrderFunc(f func(*Node)) {
	if (node == nil) {
		return
	}
	node.Left.InOrderFunc(f)
	f(node)
	node.Right.InOrderFunc(f)
}

// 使用channel实现遍历
func (node *Node) InOrderWithChannel() chan *Node {
	// 创建*Node管道并返回
	out := make(chan *Node)
	// 新建gorountine逻辑，遍历并向管道发送数据
	go func() {
		node.InOrderFunc(func(node *Node) {
			out <- node
		})
		// 数据发送完成后close管道(一般是发送者执行close)
		close(out)
	}()
	return out
}