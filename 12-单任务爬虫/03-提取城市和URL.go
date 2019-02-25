package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
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

	// 获取城市列表
	getAllCityList(bytes)

}

func getAllCityList(html []byte) {
	// 使用[^X]匹配到非X的表达式，提取出url和城市名
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)" data-v-[0-9a-zA-Z]+[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(html, -1)
	for _, matchStrs := range matches {
		url := string(matchStrs[1])
		city := string(matchStrs[2])
		fmt.Println(city, " ", url)
		fmt.Println()
	}
	fmt.Println(len(matches))
}