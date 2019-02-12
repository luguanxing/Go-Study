package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// 自定义出现panic的逻辑测试
func paincTestHandler(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

// 自定义出现UserError的逻辑测试
type testingControllerUserError string

func (e testingControllerUserError) Error() string {
	return e.Message()
}

func (e testingControllerUserError) Message() string {
	return string(e)
}

func userErrorTestHandler(writer http.ResponseWriter, request *http.Request) error {
	return testingControllerUserError("Testing User Error")
}

// 自定义出现System错误的逻辑测试
func errNotFoundTestHandler(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNoPermissonTestHandler(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknownTestHandler(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("Unknown Error")
}

// 自定义无错误的逻辑测试
func noErrorTestHandler(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprint(writer, "没有错误啦")
	return nil
}

// 模拟出错情况样例
var tests = []struct {
	handler webHandler // 要包装的逻辑，会出现各种错误
	code    int        // 出现错误应返回的httpcode
	message string     // 出现错误应显示的信息
}{
	{paincTestHandler, 500, http.StatusText(500)},
	{userErrorTestHandler, 400, "Testing User Error"},
	{errNotFoundTestHandler, 404, http.StatusText(404)},
	{errNoPermissonTestHandler, 403, http.StatusText(403)},
	{errUnknownTestHandler, 500, http.StatusText(500)},
	{noErrorTestHandler, 200, "没有错误啦"},
}

// 调函数得结果测试
func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		// 包装目标逻辑函数，进行测试
		testHandler := errWrapper(tt.handler)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.google.com", nil)
		// 测试模拟请求
		testHandler(response, request)
		// 检查模拟结果
		bytes, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(bytes), "\n")
		if response.Code != tt.code || body != tt.message {
			t.Errorf("错误：应为(%d, %s)，实为(%d, %s)", tt.code, tt.message, response.Code, body)
		}
	}
}

// 在服务器中进行测试
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		testHandler := errWrapper(tt.handler)
		server := httptest.NewServer(http.HandlerFunc(testHandler))
		response, _ := http.Get(server.URL)
		// 检查模拟结果
		bytes, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(bytes), "\n")
		if response.StatusCode != tt.code || body != tt.message {
			t.Errorf("错误：应为(%d, %s)，实为(%d, %s)", tt.code, tt.message, response.StatusCode, body)
		}
	}
}
