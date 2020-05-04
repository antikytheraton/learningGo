/*
	STRUCTS MANIPULATION
*/

package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	episodes   []string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "John",
		episodes:  []string{},
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}
	fmt.Println(aDoctor)

	bDoctor := struct{ name string }{name: "John Pertwee"}
	fmt.Println(bDoctor)

	cDoctor := bDoctor
	cDoctor.name = "Tom Baker"
	fmt.Println(bDoctor, cDoctor)

	dDoctor := &bDoctor
	dDoctor.name = "Jesse James"
	fmt.Println(bDoctor, dDoctor)
}
