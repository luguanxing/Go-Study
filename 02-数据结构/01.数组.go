package main

import (
	"fmt"
	"strconv"
)

// 数组作为参数传递时会拷贝数组(函数内部使用的是副本)
func printArray(arr [5]int) {
	// 使用range遍历数组同时取序号i和值val
	for i, val := range arr {
		fmt.Println("arr[" + strconv.Itoa(i) + "]=" + strconv.Itoa(val))
	}
	fmt.Println()
}

// 指针修改数组
func modifyArray(arr *[5]int) {
	arr[0] = 100
}

func main() {
	var arr1 [5]int                  //使用var定义数组，无需设置初始值
	arr2 := [3]int{1, 3, 5}          //使用:=定义数组，必须指定设置初始值
	arr3 := [...]int{2, 4, 6, 8, 10} //省略个数，有编译器根据值自动设定长度
	var grid [4][5]int               //二维数组
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	fmt.Println()

	printArray(arr3)
	modifyArray(&arr3) //和C语言不同，数组名本身不可作为地址，需要加&
	printArray(arr3)
}
