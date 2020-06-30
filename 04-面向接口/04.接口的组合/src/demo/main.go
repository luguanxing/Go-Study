package main

import (
	"demo/service"
	"demo/service/impl"
	"fmt"
)

// 使用组合接口
func testSession(rp service.RetrieverPoster)  {
	fmt.Println("test post & get...")
	get := rp.Get("www.baidu.com")
	post := rp.Post("www.baidu.com", nil)
	fmt.Println(get)
	fmt.Println(post)
}

func main() {
	var rp service.RetrieverPoster
	rp = &impl.RetrieverPosterImpl{}
	testSession(rp)
}
