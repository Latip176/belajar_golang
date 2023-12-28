package main

import "fmt"

func edit_data(key string, value string, data map[string]string) map[string]string {
	data[key] = value
	return data
}

func main() {
	myData := map[string]string{
		"nama":   "Latif harkat",
		"alamat": "Sukabumi, Jawa Barat.",
	}
	edit_data("alamat", "Indonesia", myData)

	fmt.Println(myData)
}
