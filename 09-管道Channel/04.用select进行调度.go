package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 创建管道和对应的发送逻辑
func makeIntChannel() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	c1, c2 := makeIntChannel(), makeIntChannel()
	// 总时间超过限制则关闭
	closeTimeChannel := time.After(10 * time.Second)
	// 在管道中谁数据先来select先执行其对应的逻辑
	for {
		select {
		case n := <-c1:
			go func() {
				fmt.Println("c1:", n)
			}()
		case n := <-c2:
			go func() {
				fmt.Println("c2:", n)
			}()
		case <-time.After(800 * time.Millisecond): // 每两次select间隔太长则执行
			go func() {
				fmt.Println("no select action in 800ms")
			}()
		case <-closeTimeChannel:
			fmt.Println("select loop stop")
			return
		}
	}
}
