package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer()  {
	// defer能保证在return或panic使函数返回前能执行，顺序是先进后出(栈)，一般用于成对对应的操作
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	// 模拟函数执行中断
	panic("break")
}

func writeFile(filename string)  {
	// 建立文件create对应关闭操作close
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 建立writer对应flush
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 1; i <= 10; i++ {
		fmt.Fprintf(writer, "%d\n", i)
	}
}

func main() {
	writeFile("hello.txt")
}
