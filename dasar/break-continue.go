package main

import "fmt"

func main() {
	type kalimat string

	slice := []kalimat{
		"muhammad",
		"latif",
		"harkat",
	}
	data := map[kalimat]kalimat{
		"firstName": "muhammad",
		"center":    "latif",
		"lastName":  "harkat",
	}

	for i := 0; i < len(slice); i++ {
		for _, value := range data {
			if slice[i] == "latif" && value == "latif" {
				fmt.Println("Found!", value)
				break
			}
		}
	}
	for _, value := range data {
		if value == "latif" {
			continue
		}
		fmt.Println(value)
	}
}
