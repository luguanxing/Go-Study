package main

import (
	"fmt"
	"time"
)

// goroutine定义：函数加上go可给调度器运行、无需区分异步函数、调度器在合适时切换，使用-race检测数据访问冲突
// goroutine可能切换的点：I/O，select，channel，等待锁，函数调用(有时)，runtime.Gosched()

func main() {
	// 1000个协程映射到实际的核数个线程内执行
	for i := 0; i < 1000; i++ {
		// 使用协程并发执行函数(匿名)，并传参保证内部的i隔离
		go func(i int) {
			for {
				fmt.Println("[", i, "]")
			}
		}(i)
	}
	// 延迟main退出时间，使打印来得及
	time.Sleep(time.Minute)
}
