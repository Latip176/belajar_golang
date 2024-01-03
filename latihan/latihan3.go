package main

import "fmt"

/**
Latihan 3 Slice dan Manipulasi Data serta Pengaksesan data pada Slice

*/

func cek_login(username string, password string) bool {
	if username == "admin" && password == "Indonesia" {
		return true
	} else {
		return false
	}
}

func main() {
	data := []interface{}{
		map[string]string{
			"username": "admin",
			"password": "Indonesia",
		},
		[...]string{
			"Apel",
			"Mangga",
			"Manggis",
			"Kelengkeng",
		},
	}

	fmt.Println(data)

	username := data[0].(map[string]string)["username"]
	password := data[0].(map[string]string)["password"]

	if cek_login(username, password) {
		fmt.Println("Login Berhasil!")
	} else {
		fmt.Println("Login gagal!")
	}

	fmt.Println(data[1].([4]string))

	// add data
	data = append(data, "hayang modol")
	fmt.Println(data)
}
