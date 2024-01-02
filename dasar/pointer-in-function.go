package main

import "fmt"

type User struct {
	Username, Password string
}

func changeData(wht string, to string, data *User /* use '*' untuk penanda pointer */) interface{} {
	switch wht {
	case "username":
		data.Username = to
	case "password":
		data.Password = to
	}

	return data
}

func main() {
	data := User{Username: "admin", Password: "latifharkat"}
	fmt.Println(data)

	// menggunakan '&' untuk pointer
	changeData("username", "latip176", &data)
	changeData("password", "gayung", &data)

	fmt.Println(data)
}
