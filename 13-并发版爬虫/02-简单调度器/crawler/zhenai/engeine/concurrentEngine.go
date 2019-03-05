package engine

import (
	"crawler/zhenai/types"
	"log"
)

// 并发版执行引擎，包含调度器和执行者数
type ConcurrentEngine struct {
	Scheduler   Scheduler // 任务调度器(队列、管道)
	WorkerCount int       // 工作线程数
}

type Scheduler interface {
	Submit(request types.Request) // 提交任务
	InitChan(chan types.Request)  //初始化管道
}

func (c *ConcurrentEngine) Run(seeds ...types.Request) {
	// 建立输入*1，输出*WorkerCount管道
	in := make(chan types.Request)
	out := make(chan types.ParseResult)
	c.Scheduler.InitChan(in)
	for i := 0; i < c.WorkerCount; i++ {
		createWorker(in, out)
	}
	// 初始化任务队列
	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}
	// 获得解析后的更多任务结果，加入队列中
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("item ：", item)
		}
		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

// 输入输出管道对接逻辑，使用协程go func，获取页面解析结果输出
func createWorker(in chan types.Request, out chan types.ParseResult) {
	go func() {
		for {
			request := <-in // 没有人接收时死锁
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parseResult // 与上方对应，都在输出时死锁
		}
	}()
}
