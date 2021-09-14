package main

import "fmt"

func mul(a, b int) int {

	res := a * b
	fmt.Println("result:", res)
	return 0
}

func show() {
	fmt.Println("hello", "world")
}

func main() {

	mul(10, 20)
	defer mul(12, 20) //Last In First Out
	defer mul(5, 5)
	show()
}
