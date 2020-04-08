package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Aaron", age: 29}
	fmt.Println(p)
	fmt.Println(p.name)
}
