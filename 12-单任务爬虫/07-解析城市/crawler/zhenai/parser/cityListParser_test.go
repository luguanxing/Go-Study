package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	// 使用保存的文件，避免直接抓取数据失败
	bytes, err := ioutil.ReadFile("test_data_city_list.html")
	if err != nil {
		panic(err)
	}
	// 测试解析
	result := ParseCityList(bytes)
	const resultSize = 470
	var expectedUrls = []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	var expectedCitys = []string{
		"阿坝",
		"阿克苏",
		"阿拉善盟",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("Requests长度应为%d，实际为%d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("第%d个URL应为%s，实际为%s", i, result.Requests[i].Url, url)
		}
	}
	if len(result.Items) != resultSize {
		t.Errorf("Items长度应为%d，实际为%d", resultSize, len(result.Requests))
	}
	for i, city := range expectedCitys {
		if result.Items[i].(string) != city {
			t.Errorf("第%d个city应为%s，实际为%s", i, result.Items[i].(string), city)
		}
	}
}
