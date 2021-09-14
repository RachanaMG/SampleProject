package main

import "fmt"

func main() {

	odd := [7]int{1, 3, 5, 7, 9, 11, 13}

	for _, item := range odd {

		fmt.Println(item)
	}
}
