package parser

import (
	"crawler/zhenai/types"
	"regexp"
)

// 解析城市URL和名称的正则表达式
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)" data-v-[0-9a-zA-Z]+[^>]*>([^<]+)</a>`

// 解析城市html返回解析结果
func ParseCityList(bytes []byte) types.ParseResult {
	result := types.ParseResult{}
	// 使用[^X]匹配到非X的表达式，提取出url和城市名
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(bytes, -1)
	for _, matchStrs := range matches {
		url := string(matchStrs[1])
		cityName := string(matchStrs[2])
		// 返回城市任务(url和对应解析器)
		result.Requests = append(result.Requests, types.Request{
			Url:       string(url),
			ParseFunc: types.NilParser,
		})
		result.Items = append(result.Items, cityName);
	}
	return result
}
