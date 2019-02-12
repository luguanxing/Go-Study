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
		// 保护出错防止panic到系统，保护意料之内的Error
		defer func() {
			r := recover()
			if r != nil {
				log.Printf("Panic出错保护:%v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		// 使用自己的业务处理逻辑，如果抛出错误则当场处理，报错内容展示给用户
		err := handler(writer, request)
		if err != nil {
			// 日志输出错误
			log.Printf("请求出错:%s", err.Error())
			// 判断是否UserError，直接展示报错内容
			userErr, isUserErr := err.(controller.UserError)
			if (isUserErr) {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			} else {
				// 细分处理具体System错误类型并进行对应处理，报错内容不展示给用户
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
}

func main() {
	http.HandleFunc("/", errWrapper(controller.FileListController))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
