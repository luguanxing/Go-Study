package main

import (
	"crawler/zhenai/engeine"
	"crawler/zhenai/parser"
	"crawler/zhenai/scheduler"
	"crawler/zhenai/types"
)

func main() {
	e := engine.ConcurrentEngine{
		&scheduler.SimpleScheduler{},
		100,
	}
	e.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
