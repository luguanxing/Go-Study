package main

import "fmt"

func printSliceLenAndCap(s []int) {
	fmt.Println("len=", len(s), " cap=", cap(s))
}

func testAppend() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)
	s1 := arr[2:6]
	fmt.Println(s1)
	s2 := s1[3:5]
	fmt.Println(s2)
	s3 := append(s2, 10)	// append改动切片会影响到底层数组
	fmt.Println(s3)
	s4 := append(s3, 11) // 这时底层的数组重新分配变化了
	fmt.Println(s4)
	s5 := append(s4, 12)
	fmt.Println(s5)
	fmt.Println(arr)
}

func testCreate() {
	var s []int // 初始值为nil
	for i := 0; i < 10; i++ {
		s = append(s, 2*i+1)
		fmt.Println(s)
		printSliceLenAndCap(s)
	}
	s1 := []int{2, 4, 6, 8}
	printSliceLenAndCap(s1)
	s2 := make([]int, 16)
	printSliceLenAndCap(s2)
	s3 := make([]int, 16, 32)
	printSliceLenAndCap(s3)
}

func testCopy() {
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, len(s1)) // 需要预先分配空间
	copy(s2, s1)
	fmt.Println(s1)
	printSliceLenAndCap(s1)
	fmt.Println(s2)
	printSliceLenAndCap(s2)
}

func testDelete() {
	s1 := []int{2, 4, 6, 8}
	s2 := append(s1[0:1], s1[2:]...) // 删除s1[1]即在s1[0:1]添加s1[2:]以及可变元素...
	fmt.Println(s2)
	printSliceLenAndCap(s2)
	s3 := s2[1:]	// 删除头尾元素
	s3 = s3[:len(s3)-1]
	fmt.Println(s3)
	printSliceLenAndCap(s3)
}

func main() {
	testAppend()
	testCreate()
	testCopy()
	testDelete()
}
