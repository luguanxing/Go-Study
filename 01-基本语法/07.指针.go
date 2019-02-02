package main

import "fmt"

func swap(a, b int) {
	b, a = a, b
}

func swapPoniter(a, b *int) {
	*b, *a = *a, *b
}

func swapNoPoniter(a, b int) (int, int){
	return b, a
}

func main() {
	a, b := 1, 2
	swap(a, b)
	fmt.Println(a, b)

	swapPoniter(&a, &b)
	fmt.Println(a, b)

	a, b = swapNoPoniter(a, b)
	fmt.Println(a, b)
}
