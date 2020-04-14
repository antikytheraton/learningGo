package main

import "fmt"

func main() {
	func() {
		fmt.Println("hello go!")
	}()

	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	f := func() {
		fmt.Println("Hi Go!")
	}
	fmt.Println(&f)
	f()
}
