package main

import "fmt"

func main() {
	type angka int

	var x angka = 10
	var y angka = 20

	var hasil = x + y

	fmt.Println("Hasil dari pertambahan", x, "+", y, "=", hasil)

	x += 10 // augment operator
	fmt.Println(x)

	y++ // unary operator
	fmt.Println(y)

}
