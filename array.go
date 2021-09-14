package main

import "fmt"

func main() {

	arr := [4]string{"This", "is", "my", "class"}

	fmt.Println("Elements of the array:")

	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}

}
