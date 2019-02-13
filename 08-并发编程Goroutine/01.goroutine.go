package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		// 使用协程并发执行函数(匿名)，并传参保证内部的i隔离
		go func(i int) {
			for {
				a[i]++
				// 手动交出控制权，防止无限占用，使别的协程有机会运行(较公平)，一般也可不用
				runtime.Gosched()
			}
		}(i)
	}
	// 延迟main退出时间，使打印来得及
	time.Sleep(time.Millisecond)
	// 打印结果
	fmt.Println(a)
}
