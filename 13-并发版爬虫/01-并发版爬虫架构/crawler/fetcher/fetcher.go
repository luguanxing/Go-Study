package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

func Fetch(url string) ([]byte, error) {
	// 打开网页
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 判断状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong StatusCode : %d", resp.StatusCode)
	}
	// 打印内容
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func ClientFetch(url string) ([]byte, error) {
	// 设置请求信息
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	// 模拟客户端获取url请求
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 读取请求内容
	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}