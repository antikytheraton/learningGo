package main

import "fmt"

// Stringer interface
type Stringer interface {
	String() string
}

// Article type
type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("The %q book was written by %s.", b.Title, b.Author)
}

func main() {
	a := Article{
		Title:  "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}
	Print(a)
	b := Book{
		Title:  "All about Go",
		Author: "Jenny Dolphin",
		Pages:  25,
	}
	Print(b)
}

// Print is a printer function
func Print(a Stringer) {
	fmt.Println(a.String())
}
