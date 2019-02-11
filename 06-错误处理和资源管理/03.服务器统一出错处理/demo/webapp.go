package main

import (
	"demo/controller"
	"log"
	"net/http"
	"os"
)

// 自己的业务处理逻辑，能抛出error给上级处理
type webHandler func(writer http.ResponseWriter, request *http.Request) error

// 上级包装函数，专门用来统一处理抛出的error
func errWrapper(handler webHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 使用自己的业务处理逻辑，如果抛出错误则当场处理
		err := handler(writer, request)
		if err != nil {
			// 日志输出错误
			log.Printf("请求出错:%s", err.Error())
			// 细分处理具体错误类型并进行对应处理
			httpCode := http.StatusOK
			switch {
			case os.IsNotExist(err):
				httpCode = http.StatusNotFound
			case os.IsPermission(err):
				httpCode = http.StatusForbidden
			default:
				httpCode = http.StatusInternalServerError
			}
			// 显示包装的原因
			http.Error(writer, http.StatusText(httpCode), httpCode)
		}
	}
}

func main() {
	http.HandleFunc("/static/", errWrapper(controller.FileListController))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
