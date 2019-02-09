package main

import (
	"demo/retriever"
	"demo/retriever/impl"
	"fmt"
)

func getBaiduIndex(r retriever.Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r retriever.Retriever
	fmt.Printf("%T %v %p\n", r, r, r)
	r = &impl.RetrieverImpl{}
	fmt.Printf("%T %v %p\n", r, r, r)
	// 接口变量 = 实现者类型 + 实现者指针，所以接口变量几乎不需要使用接口的指针
	rr, ok := r.(*impl.RetrieverImpl)
	if ok {
		fmt.Printf("取得接口里真正的实现类 %p\n", rr)
	}
}
