package main

import "fmt"

type number int

func addData(data ...number) (data_array []number) {
	for _, i := range data {
		num := i + 10
		data_array = append(data_array, num)
		fmt.Printf("%d + %d = %d\n", i, 10, num)
	}
	return data_array
}

func main() {
	myData := addData(10, 20, 30, 50, 60)
	fmt.Println(myData)
}
