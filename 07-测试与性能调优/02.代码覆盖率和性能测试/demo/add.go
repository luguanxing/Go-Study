package test

import "fmt"

func add(a, b int) int {
	return a + b
}

// 代码生成覆盖率文件go test -coverprofile=c.out
// 图形化界面展示go tool cover -html=c.out
func usesless() {
	// 代码覆盖率跑不到
	fmt.Println("Unreachable")
}