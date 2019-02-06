package queue

// 把MyQueue作为切片[]int的别名
type MyQueue []int

func (q *MyQueue) Push(n int) {
	// 修改了自己的切片slice
	*q = append(*q, n)
}

func (q *MyQueue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *MyQueue) IsEmpty() bool {
	return len(*q) == 0
}
