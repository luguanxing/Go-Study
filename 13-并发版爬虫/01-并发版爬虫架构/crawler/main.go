package main

import (
	"crawler/zhenai/engeine"
	"crawler/zhenai/parser"
	"crawler/zhenai/types"
)

func main() {
	//engine.Run(types.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})
	engine.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun/aba",
		ParseFunc: parser.ParseCity,
	})
}
