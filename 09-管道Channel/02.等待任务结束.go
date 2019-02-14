package main

import (
	"fmt"
	"sync"
)

func createReceiver(i int) (chan<- int, chan bool) {
	c := make(chan int)
	done := make(chan bool)
	go func() {
		for {
			fmt.Println("channels[", i, "] -> ", <-c)
			// 打印完毕，使用管道done通知完成(阻塞式，要被取后才能继续，否则使用gorountine)
			done <- true
		}
	}()
	return c, done
}

func channelDemo() {
	var channels [10]chan<- int
	var dones [10] chan bool
	for i := 0; i < 10; i++ {
		channels[i], dones[i] = createReceiver(i)
	}
	for i := 0; i < 10; i++ {
		// 放入数据
		channels[i] <- i
	}
	for i := 0; i < 10; i++ {
		// 等待done结果，能取出done说明已打印完
		<-dones[i]
	}
	for i := 0; i < 10; i++ {
		// 放入数据
		channels[i] <- i + 100
	}
	for i := 0; i < 10; i++ {
		// 等待done结果，能取出done说明已打印完
		<-dones[i]
	}
}

func createReceiverWg(i int, done func()) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Println("channels[", i, "] -> ", <-c)
			done()	// 打印完执行完成逻辑
		}
	}()
	return c
}

func channelDemoWg() {
	var channels [10]chan<- int
	var wg sync.WaitGroup	// 任务组管理器
	for i := 0; i < 10; i++ {
		channels[i] = createReceiverWg(i, func() {
			wg.Done()	// 完成一个任务的逻辑
		})
	}
	// 添加数个任务(建议不要在循环中wg.Add(1)，因为可能在此前执行了任务逻辑和wg.Done()使变成负数)
	wg.Add(20)
	for i := 0; i < 10; i++ {
		channels[i] <- i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- i + 10000
	}
	wg.Wait()	// 等待所有任务完成
}

func main() {
	//channelDemo()
	channelDemoWg()
}
