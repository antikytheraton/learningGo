package main

import (
	"fmt"
)

func main() {
 var w Writer = ConsoleWriter{}
 n, err := w.Write([]byte("Hello World!"))
 fmt.Println(n, err)
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {}

func(cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}