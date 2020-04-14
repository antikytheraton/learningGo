package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("something went wrong")
	fmt.Println("end")
}
