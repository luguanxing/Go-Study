package main

import (
	"fmt"
	"reflect"
	"runtime"
)

// 函数可以有多个返回值，获取后不需要的可以填下划线_
func div(a, b int) (q, r int, err error) {
	if (b == 0) {
		return q, r, fmt.Errorf("被除数为0")
	}
	q = a / b
	r = a % b
	return q, r, nil
}

// 函数式编程
func apply(opFunc func(int, int) int, a, b int) int {
	p := reflect.ValueOf(opFunc).Pointer()
	opFuncName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling %s with args : %d, %d\n", opFuncName, a, b)
	return opFunc(a, b)
}

func add(a, b int) int {
	return a + b
}

// 可变参数列表
func sum(nums ...int) int {
	s := 0
	for i := range nums {
		s += nums[i]
	}
	return s
}

func main() {
	// 自动生成返回值快捷键Ctrl+Alt+V
	q, r, err := div(13, 0)
	fmt.Println(q, r, err)
	fmt.Println(apply(add, 1, 2))
	fmt.Println(apply(func(a int, b int) int {
		//匿名函数
		return a*a + b*b
	}, 1, 2))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
}
