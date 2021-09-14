package main

import (
	"fmt"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		//	time.Sleep(500 * time.Millisecond)
		fmt.Println(str)
	}
}

func main() {
	fmt.Println("hello")

	go display("Welcome")

	go display("Golang")

	//time.Sleep(3500 * time.Millisecond)
	fmt.Println("world")
}
