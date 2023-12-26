package main

import "fmt"

func main() {
	/*
		map ini tipe data seperti dictinory atau object jika di bahasa lain.
		untuk membuat map formatnya:
		data := map[TypeKey]TypeValue
	*/
	type kalimat string
	type latif bool

	data := map[kalimat]latif{
		"laki-laki": true,
		"ganteng":   true,
		"perempuan": false,
	}
	fmt.Println(data)
	fmt.Println(data["laki-laki"])
	fmt.Println(data["perempuan"])
	fmt.Println(data["ganteng"])

	delete(data, "perempuan") // hapus data perempuan
	data["programmer"] = true // menambahkan data baru ke map
	fmt.Println(data)
}
