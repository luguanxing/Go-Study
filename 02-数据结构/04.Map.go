package main

import "fmt"

func main() {
	// 创建map
	m := map[string]string {
		"van" : "my name is van",
		"窃格瓦拉" : "打工是不可能打工的",
		"六小龄童" : "今年下半年",
	}
	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m, m2, m3)

	// 遍历map
	for k, v := range m {
		fmt.Println(k, "说过:", v, "...")
	}
	fmt.Println()

	// 取值map，并判断是否存在
	六小龄童语录 := m["六小龄童"]
	fmt.Println(六小龄童语录)
	vanword, exist := m["van"]
	fmt.Println("m['van']是否存在:",exist, "=>", vanword)
	vanword2, exist := m["van darkholme"]	//取不存在的值为空字符串
	fmt.Println("m['van darkholme']是否存在:",exist, "=>", vanword2)

	// 删除map值
	delete(m, "van")
	delete(m, "窃格瓦拉")
	fmt.Println(m)
}
