package main

import "fmt"

func ops(p, q int) (int, int, int) {
	return p - q, p * q, p + q
}

func main() {

	var v1, v2, v3 = ops(4, 2)

	fmt.Printf("Result is: %d", v1)
	fmt.Printf("\nResult is: %d", v2)
	fmt.Printf("\nResult is: %d", v3)
}
