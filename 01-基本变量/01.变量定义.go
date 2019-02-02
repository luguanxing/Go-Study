package main

import "fmt"

// 全局变量
var (
	str     = "hello world!"
	packInt = 1
)

func testVariable() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func testVariableInit() {
	var a, b int = 3, 4
	var s string = "fa q"
	fmt.Printf("%d %d %q\n", a, b, s)
}

func testVariableType() {
	var a, b, s = 1, 2, "van"
	fmt.Printf("%d %d %q\n", a, b, s)
}

func testVariableLocal() {
	a, b, s := 5, 6, "deep dark fantasty"
	fmt.Printf("%d %d %q\n", a, b, s)
}

func main() {
	fmt.Printf("%d\n", packInt)
	fmt.Print(str + "\n")
	fmt.Print("hello world\n")
	testVariable()
	testVariableInit()
	testVariableType()
	testVariableLocal()
}
