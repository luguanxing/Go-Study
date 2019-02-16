package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// 直接访问url获取返回内容
func httpGet()  {
	// 获取url请求
	resp, err := http.Get("http://www.imooc.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取请求内容
	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

// 模拟客户端使用请求访问url获取返回内容
func clientHttpGet() {
	// 设置请求信息
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	// 模拟客户端获取url请求
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取请求内容
	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

// 使用更高级可定制化的client
func clientStructGet()  {
	// 设置请求信息
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	// 定义可定制化的client
	client := http.Client{
		// 重定向逻辑
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect->req:", req)
			fmt.Println("Redirect->via:", via)
			return nil
		},
	}


	// 模拟客户端获取url请求
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取请求内容
	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func main() {
	// httpGet()
	// clientHttpGet()
	clientStructGet()
}
