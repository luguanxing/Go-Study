package scheduler

import "crawler/zhenai/types"

type SimpleScheduler struct {
	WorkerChan chan types.Request
}

func (s *SimpleScheduler) InitChan(c chan types.Request) {
	s.WorkerChan = c
}

func (s *SimpleScheduler) Submit(request types.Request) {
	// 使用go rountine避免死锁
	go func() {
		s.WorkerChan <- request
	}()
}



