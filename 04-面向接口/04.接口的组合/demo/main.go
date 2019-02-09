package main

import (
	"demo/service"
	"demo/service/impl"
	"fmt"
)

func getBaiduIndex(r service.Retriever) string {
	return r.Get("http://www.baidu.com")
}

func postData(r service.Poster) string {
	return r.Post("http://www.baidu.com", map[string]string{
		"name": "van darkholme",
		"job":  "artist",
	})
}

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
