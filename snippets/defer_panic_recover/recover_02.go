package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			// this is for studying the panic error
			// panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}
