package main

import (
	"fmt"
	dep "github.com/RichardInnocent/example-go-dependency"
)

func main() {
	fmt.Println(dep.ExampleStruct{
		Value: "Hello, world!",
	})
}
