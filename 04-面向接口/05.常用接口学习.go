package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Student struct {
	sid  string
	name string
}

// 使用系统提供（也可以说是符合定义标准即可）的stringer接口实现类似java的toString方法
func (s Student) String() string {
	return "student[sid=" + s.sid + ", name=" + s.name + "]"
}

// 使用Reader接口统一转换输入参数格式，便于输出
func printReaderContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	// 测试类的打印
	student := Student{"10000", "马云"}
	fmt.Println(student)

	// 测试格式转换输出
	testStr :=
`
	今年下半年，
	中美合拍的西游记即将正式开机，
	我继续扮演美猴王孙悟空，
	我会用美猴王艺术形象努力创造一个正能量的形象，
	文体两开花，
	弘扬中华文化，
	希望大家能多多关注。
`
	printReaderContents(strings.NewReader(testStr))
}
