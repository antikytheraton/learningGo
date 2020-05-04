package main

import (
	"fmt"
)

func main() {
	// for loop sentence
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// iterate an array
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	// iterate a dictionary
	dict := make(map[string]string)
	dict["a"] = "alpha"
	dict["b"] = "beta"
	for key, value := range dict {
		fmt.Println("key:", key, "value:", value)
	}
}
