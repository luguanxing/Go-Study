package controller

import (
	"io/ioutil"
	"net/http"
	"os"
)

func FileListController(writer http.ResponseWriter, request *http.Request) error {
	// 获取文件名
	path := request.URL.Path[len("/"):]
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		return err	// 向上抛出错误，由上级处理
	}
	defer file.Close()
	// 读取文件
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err	// 向上抛出错误，由上级处理
	}
	// 输出文件内容
	writer.Write(bytes)
	return nil
}
