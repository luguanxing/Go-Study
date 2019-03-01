package types

// 解析任务
type Request struct {
	Url       string                   //抓取的URL
	ParseFunc func([]byte) ParseResult //该URL对应的解析函数，返回解析后的结果
}

// 解析结果，返回更多任务和附加对象
type ParseResult struct {
	Requests []Request     //解析后的新任务
	Items    []interface{} //任何对象
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
