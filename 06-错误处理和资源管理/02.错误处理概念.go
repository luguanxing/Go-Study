package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func writeFile(filename string) {
	// 建立文件create对应关闭操作close
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		// 处理错误
		fmt.Println("处理文件出错 ->", err)
		// 进入源码可以看该err可以转换成什么类型
		// If there is an error, it will be of type *PathError.
		pathError := err.(*os.PathError)
		fmt.Printf("pathError.Op = %s\npathError.Path = %s\npathError.Err = %s\n",
			pathError.Op,
			pathError.Path,
			pathError.Err,
		)
		return
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
	// 测试错误处理
	writeFile("hello.txt")

	// 可建立自己的error
	err := errors.New("my error")
	fmt.Println(err)
}
