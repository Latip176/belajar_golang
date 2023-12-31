package main

import (
	"fmt"
)

func random() interface{} {
	return "Hello World!"
}

func main() {
	data := random()

	switch value := data.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unkown Type")
	}
}
