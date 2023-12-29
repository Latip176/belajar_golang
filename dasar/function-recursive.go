package main

import "fmt"

var no = 0 // global variables

func recursive(data ...string) (answer string) {
	for _, v := range data {
		if data[no] == "admin" {
			answer = "admin"
			return answer
		} else {
			no++
			fmt.Println("Nope", v)
			recursive(data...) // recursive function
		}
	}
	return
}

func main() {
	data := []string{
		"latif",
		"admin",
		"root",
	}
	fmt.Println("Ban", recursive(data...))
}
