package main

import "fmt"

type Data struct {
	Name string
}

func main() {
	data := Data{Name: "Latif Harkat"}
	data2 := &data // pointer

	data2.Name = "Omen Craft" // ubah data Name jadi Omen Craft dengan pointer: cara ke 1
	fmt.Println(data)

	var data3 *Data = &data     // pointer
	data3.Name = "Latif Harkat" // ubah data Name jadi Omen Craft dengan pointer: cara ke 2

	fmt.Println(data)
}
