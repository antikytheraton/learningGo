package main

import (
	"fmt"
)

func main() {
	//a := [5]int{}
	//a[2] = 7

	a := []int{1, 2, 3, 4, 5}
	a = append(a, 7)
	fmt.Println(a)
}
