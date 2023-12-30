package main

import "fmt"

type Data struct {
	Name, Address string
	Age, TB       int
}

func (Information Data) showData(companyname string) {
	fmt.Println("Welcome to", companyname, "Company!")
	fmt.Printf("Name: %s\nAddress: %s\nAge: %d\nTB: %d\n", Information.Name, Information.Address, Information.Age, Information.TB)
}
func (Information Data) showName() {
	fmt.Println("Name:", Information.Name)
}
func (Information Data) showAddress() {
	fmt.Println("Address:", Information.Age)
}
func (Information Data) showAge() {
	fmt.Println("Age:", Information.Address)
}
func (Information Data) showTB() {
	fmt.Println("TB:", Information.TB)
}

func main() {
	var MyData Data
	data1 := MyData
	data1.Name = "Latif Harkat"
	data1.Address = "Sukabumi, Jawa Barat"
	data1.Age = 18
	data1.TB = 175

	data1.showData("Sungut Lele")
	data1.showName()
	fmt.Println()

	data2 := Data{
		Name:    "Omen Craft",
		Address: "Tokyo, Japan",
		Age:     18,
		TB:      170,
	}
	data2.showData("Corporosion")
}
