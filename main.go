package main

import (
	"fmt"

	"github.com/AbdulConsole/Learning-Go/learn"
)

func main() {
	fmt.Println("Hello World!")

	// Example fetching from first package
	getTwo := learn.Two()
	fmt.Println(getTwo)

	// Example fetching from second package
	getThree := learn.Three()
	fmt.Println(getThree)
}
