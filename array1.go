package main

import "fmt"

func main() {

	arr := [3][3]string{{"C#", "C", "Python"},
		{"Java", "Scala", "Perl"},
		{"C++", "Go", "HTML"}}

	fmt.Println("Elements of Array 1")
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Println(arr[x][y])
		}
	}

	var arr1 [2][2]int
	arr1[0][0] = 100
	arr1[0][1] = 200
	arr1[1][0] = 300
	arr1[1][1] = 400

	fmt.Println("Elements of array 2")
	for p := 0; p < 2; p++ {
		for q := 0; q < 2; q++ {
			fmt.Println(arr1[p][q])

		}
	}
}
