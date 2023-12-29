package main

import "fmt"

func finished() {
	fmt.Println("Program Selesai")
}
func runApplication(name string) {
	defer finished() // akan di eksekusi walau jika ada error
	fmt.Println("Welcome", name)
}

func main() {
	runApplication("Latif Harkat")
}
