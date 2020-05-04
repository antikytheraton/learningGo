package main

import "fmt"

func main() {
	a := [3]int{3, 6, 9}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v %p %p\n", a, b, c)
}
