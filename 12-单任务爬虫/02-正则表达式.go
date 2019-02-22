package main

import (
	"fmt"
	"regexp"
)

const text =
`
test abc
abc@qq.com haha
sdaf
email = 123@abc.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9.]+)`)
	// 查找正则匹配所有结果
	match := re.FindAllString(text, -1)
	fmt.Println(match)
	fmt.Println()
	// 查找子匹配
	submatch := re.FindAllStringSubmatch(text, -1)
	for _, m := range submatch {
		fmt.Println(m)
	}
}
