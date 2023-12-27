package main

import "fmt"

func main() {
	/*
		perulangan di golang hanya ada for loop
	*/

	// for statment
	for i := 1; i <= 10; i++ {
		fmt.Printf("Perulangan ke %d\n", i)
	}

	// for range
	type d string

	data := []d{
		"latif",
		"fvcking",
		"harkat",
	} // data slice

	for _, value := range data {
		fmt.Println(value)
	}

	data2 := map[d]d{
		"nama": "latif harkat",
		"asal": "sukabumi",
	}
	for key, value := range data2 {
		fmt.Println(key, value)
	}
}
