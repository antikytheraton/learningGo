package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something went wrong")
	fmt.Println("end")
}
