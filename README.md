# Learning-Go
My Step-by-step process to learn Golang.

## Setup 

#### Step 1:
Unlike other programming languages, initializing the working directory is important in go, so in the root of folder make sure to run the command:

```go mod init <git-repo-addr>```

The above command automatically setup a ```go.mod``` file for us to work with. And basically initiate our project into a module.

#### Step 2:
Next you create the ```main.go``` file, and only first line set your package name, the generic one is ```package main```.

#### Step 3:
Set an entry point for the program, where the program will start execution from, this declaration is important for the go program to know where to begin execution. so then we create ```main()```, the main function is our entry point. we create it using the ```func``` keyword.
### NOTE - You can have only one entry point, that is one main function.

#### Step 4:
Now we need a package to output our text, for this purpose we are going to import the ```fmt``` package. 

```import "fmt" ```

Now we have a script like this:

```go
package main

import "fmt"

func main() {
	fmt.Print("Hello World!")
}
```

#### Step 5:
Running the program, you use the command. ```go run main.go``` And we have the output: ```Hello World!```

## Further Programs:
Will be in the learn package