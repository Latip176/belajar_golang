package main

import "fmt"

type Data struct {
	Name string
}
type Data2 struct {
	Name string
}
type Information interface {
	GetData() string
}

// implementasi untuk interfafe
func (data Data) GetData() string {
	return data.Name
}
func (data Data2) GetData() string {
	return data.Name
}
func showData(data Information) {
	fmt.Println("My Name:", data.GetData())
}

func main() {
	info := Data{Name: "Latif Harkat"}
	info2 := Data2{Name: "Omen Craft"}
	showData(info)
	showData(info2)
}
