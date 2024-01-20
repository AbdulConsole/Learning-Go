package main

import (
	"fmt"

	"github.com/AbdulConsole/Learning-Go/learn"
)

func main() {
	// fmt.Println("Hello World!")

	// Lesson One from first package
	getOne := learn.One()
	fmt.Print(getOne)

	// Example fetching from second package
	getTwo := learn.Two()
	fmt.Printf("\n\n\nUp Next: %s ...", getTwo)
}
