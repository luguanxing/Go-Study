package main

import "fmt"

// 函数是一等公民：参数、变量、返回值都可以是函数

// 闭包 > 函数体 = 局部变量 + 自由变量
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
}
