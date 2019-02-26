package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
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