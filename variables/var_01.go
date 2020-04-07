package main

import "fmt"

func main() {
	var i int
	i = 42

	var j int = 27

	k := 99

	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)
}
