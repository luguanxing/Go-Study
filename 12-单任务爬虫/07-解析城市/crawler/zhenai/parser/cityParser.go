package parser

import (
	"crawler/zhenai/types"
	"regexp"
)

// <th><a href="http://album.zhenai.com/u/1633617220" target="_blank">游天</a></th>
const cityRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`

// 解析城市返回用户列表结果
func ParseCity(bytes[] byte) types.ParseResult {
	result := types.ParseResult{}
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(bytes, -1)
	for _, matchStrs := range matches {
		url := string(matchStrs[1])
		Name := string(matchStrs[2])
		// 返回用户任务(url和对应解析器)
		result.Requests = append(result.Requests, types.Request{
			Url:       string(url),
			ParseFunc: types.NilParser,
		})
		result.Items = append(result.Items, Name);
	}
	return result
}