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
	r = impl.RetrieverImpl{}
	fmt.Println(getBaiduIndex(r))
}
