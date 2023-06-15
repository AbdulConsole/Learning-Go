/*

Assignment 2
Initialize and assign a value to an integer, float, and string
variable. then print each one with a sentence on a new line
describing variable type.

*/

package main

import "fmt"

func main() {
        // Initialize and assignment of variables
        var intValue int = 42
        var floatValue float64 = 3.14
        var stringValue string = "Hello, World!"

        //Printing variables
        fmt.Println(intValue," is an integer!")
        fmt.Println(floatValue," is a float!")
        fmt.Println(stringValue,"is a string!")
}