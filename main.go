package main

import (
	"fmt"

	"github.com/AbdulConsole/Learning-Go/learn"
)

func main() {
	fmt.Println("Hello World!")

	getTwo := learn.Two()
	fmt.Println(getTwo)

	getThree := learn.Three()
	fmt.Println(getThree)
}
