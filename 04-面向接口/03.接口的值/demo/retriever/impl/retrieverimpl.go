package impl

import (
	"net/http"
	"net/http/httputil"
)

type RetrieverImpl struct {
}

// 只要实现了接口里的方法就认为实现了方法，无须耦合地指定Retriever接口
// 方法指针接收：在改变内容、结构过大、一致性时考虑使用
func (r *RetrieverImpl) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return (string(result))
}
