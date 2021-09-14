package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Name       string
	Occupation string
	Born       string
}

func main() {

	filename, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}

	defer filename.Close()

	data, err := ioutil.ReadAll(filename)

	if err != nil {
		fmt.Println(err)
	}

	var result []User

	jsonErr := json.Unmarshal(data, &result)

	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	fmt.Println(result)
}
