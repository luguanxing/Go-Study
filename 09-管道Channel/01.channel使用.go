package main

import (
	"fmt"
	"time"
)

// 创建监听channel并返回只收不发channel
func createReceiver(i int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Println("channels[", i, "] -> ", <- c)
		}
	}()
	return c
}

// 向channel发数据
func sender(n int, c chan<- int)  {
	c <- n
}

func channelDemo() {
	// 定义多个只收不发channel，需要使用多个协程才能能使用channel，否则死锁
	var channels [10]chan<- int
	for i := 0;  i < 10; i++ {
		// 向每个channel接收数据，收数据要在发数据前避免死锁
		channels[i] = createReceiver(i)
	}
	for i := 0;  i < 10; i++ {
		// 向每个channel编号并发送数据
		go sender(i, channels[i])
		go sender(i + 10, channels[i])
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel()  {
	// 建立带缓冲区的channel
	c := make(chan int, 3)
	go func(c chan int) {
		for {
			fmt.Println("channel -> ", <- c)
		}
	}(c)
	// 向缓冲区发数据，超出缓冲区没人接收则死锁
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	time.Sleep(time.Millisecond)
}

func closedChannel()  {
	c := make(chan int)
	go func(c chan int) {
		for n := range c {
			fmt.Println("channel -> ", n)
		}
	}(c)
	c <- 1
	c <- 2
	// 发送方发完数据close
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	// channelDemo()
	// bufferedChannel()
	closedChannel()
}
