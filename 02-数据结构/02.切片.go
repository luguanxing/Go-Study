package main

import "fmt"

func modifySlice(arr []int) {
	arr[0] = 100
}

func printArray(arr []int) {
	for i, val := range arr {
		fmt.Println("arr[", i, "] = ", val)
	}
	fmt.Println()
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// 切片slice，左闭右开
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[2:]=", arr[2:])
	fmt.Println("arr[:6]=", arr[:6])
	fmt.Println("arr[:]=", arr[:])
	fmt.Println()

	// 修改slice
	modifySlice(arr[2:6]);
	fmt.Println(arr)
	modifySlice(arr[:])
	fmt.Println(arr)
	modifySlice(arr[:][3:]) // 多次进行slice操作
	fmt.Println(arr)
	fmt.Println()

	printArray(arr[:])

	// slice可以扩展访问底层数组，slice=指针ptr + 长度len + 总容量cap，但只能向后扩展不能向前扩展
	arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	printArray(arr[2:6][3:5])
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("len(arr[2:6])=", len(arr[2:6]))
	fmt.Println("cap(arr[2:6])=", cap(arr[2:6]))
	fmt.Println()
	fmt.Println("arr[2:6][3:5]=", arr[2:6][3:5])
	fmt.Println("len(arr[2:6][3:5])=", len(arr[2:6][3:5]))
	fmt.Println("cap(arr[2:6][3:5])=", cap(arr[2:6][3:5]))
}
