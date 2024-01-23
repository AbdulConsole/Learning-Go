package main

import (
        "fmt"

        "github.com/AbdulConsole/Learning-Go/learn"
)

func main() {
        // fmt.Println("Hello World!")

        // Lesson One from first package
        getOne := learn.One()
        fmt.Println(getOne)
        fmt.Println("------------_-------------")

        // Example fetching from second package
        //Login in on termux
        getTwo := learn.Two()
        fmt.Printf("\n%s\n", getTwo)
        fmt.Println("------------_-------------")
}
