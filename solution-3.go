/*

Assignment 3
Prompt the user to input their first and last name
and then print them a welcome message. try storing the input
as a string variable's value.

Easy Mode: use fmt.Scanln()

Extra Credit: use fmt.Scanf()

*/

package main

import "fmt"

func main() {
        var firstName string
        var lastName string

        fmt.Print("Enter your first name: ")
        _, err := fmt.Scanf("%s \n", &firstName)
        if err != nil {
                fmt.Println(err)
                return
        }

        fmt.Print("Enter your last name: ")
        _, err = fmt.Scanf("%s \n", &lastName)
        if err != nil {
                fmt.Println(err)
                return
        }

        fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}