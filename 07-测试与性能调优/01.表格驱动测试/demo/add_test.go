package test

import "testing"

// 表格驱动测试：分离数据和测试逻辑(使用循环)，批量测试不中断，容易找出具体出错地方
// 注意要以大写Test开头
func TestAdd(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{1, 2, 3},
		{1, 1, 1},
		{1, -1, 1},
	}
	for _, tt := range tests {
		calcResult := add(tt.a, tt.b)
		if calcResult != tt.c {
			t.Errorf("add(%d, %d)得到了错误结果%d，正确结果为%d", tt.a, tt.b, calcResult, tt.c)
		}
	}
}
