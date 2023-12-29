package main

import "fmt"

func end() {
	errmessage := recover() // get pesan panic
	if errmessage != nil {
		fmt.Println("Terjadi Error. Pesan:", errmessage)
	}
	fmt.Println("Program Finished")
}

func run(err bool) {
	defer end()
	if err {
		panic("ERROR True!") // set panic error
	}
	fmt.Println("Program Runinng...")
}

func main() {
	run(false)
	run(true)
}
