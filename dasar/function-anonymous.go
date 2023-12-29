package main

import "fmt"

type blacklist func(string) bool

func login(name string, block blacklist) {
	if block(name) {
		fmt.Println("You are Blocked", name)
	} else {
		fmt.Println("Login Success! Welcome", name)
	}
}

func main() {
	ban := func(name string) bool {
		return name == "omen"
	}
	login("omen", ban)
	login("latif", func(name string) bool {
		return name == "omen"
	})
}
