package main

import "fmt"

func cekData(data string) interface{} {
	if data == "latif" {
		return 1
	} else if data == "omen" {
		return true
	} else {
		return nil
	}
}

func main() {
	data := cekData("latif")
	fmt.Println(data)
}
