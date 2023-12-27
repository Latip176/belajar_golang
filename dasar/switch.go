package main

import "fmt"

func main() {
	type kalimat string

	var perintah kalimat = "duduk"

	switch perintah {
	case "duduk":
		fmt.Println("robot duduk")
	case "salto":
		fmt.Println("robot salto")
	default:
		fmt.Println("robot tidak mengerti")
	}

	const command kalimat = "cuaca"
	cuaca_saat_ini := "hujan"

	switch command {
	case "cuaca":
		fmt.Printf("Cuaca saat ini %s", cuaca_saat_ini)
		if cuaca_saat_ini == "hujan" {
			fmt.Println(" Sediakan payung!")
		} else {
			fmt.Println(" Cuaca normal saja.")
		}
	default:
		fmt.Println("perintah tidak diketahui")
	}

	/*
		Switch di golang bisa tanpa Kondisi.
	*/
	nama := "Latif Harkat"

	switch {
	case nama == "Latif Harkat":
		fmt.Println("Selamat Datang Admin")
	default:
		fmt.Println("Anda tidak dikenal.")
	}
}
