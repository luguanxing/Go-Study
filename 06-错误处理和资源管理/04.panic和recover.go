package main

import (
	"errors"
	"fmt"
)

// 测试自定义Error错误
func testMyError() {
	panic(errors.New("My Error"))
}

// 测试计算Error
func testCalcError()  {
	b := 0
	a := 5 / b
	fmt.Println(a)
}

//  未知panic内容
func testUnknownPanic()  {
	panic(123)
}

func tryRecover(testFunc func()) {
	defer func() {
		// 调用错误处理函数
		r := recover()
		if (r != nil) {
			err, isErr := r.(error)
			if isErr {
				fmt.Println("出现Error错误 ->", err)
			} else {
				fmt.Println("出现未识别的panic内容", r)
			}
		}
	}()
	testFunc()
}

// 测试正常执行的函数
func testNormal()  {
	fmt.Println("hello world")
}

func main() {
	tryRecover(testMyError)
	tryRecover(testCalcError)
	tryRecover(testUnknownPanic)
	tryRecover(testNormal)
}
