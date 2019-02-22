package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 打开网页
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 判断状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("resp.StatusCode ==", resp.StatusCode)
	}

	// 打印内容
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}
