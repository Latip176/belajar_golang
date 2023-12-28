package main

import "fmt"

type kalimat string

func returnMultiple(data_slice []kalimat, data_map map[kalimat]kalimat) ([]kalimat, map[kalimat]kalimat) {
	return data_slice, data_map
}

func main() {
	data_slice := []kalimat{
		"mobil",
		"motor",
		"sepeda",
	}
	data_map := map[kalimat]kalimat{
		"data1": "mobil",
		"data2": "motor",
		"data3": "sepeda",
	}

	slice, mapp := returnMultiple(data_slice, data_map)
	fmt.Println(slice)
	fmt.Println(mapp)
}
