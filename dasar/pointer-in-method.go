package main

import "fmt"

/**
pointer in method lebih reccomended, karena tidak terlalu makan ram.
*/

type Cat struct {
	Name string
	Age  int
}

func (c *Cat /* pake '*' untuk penanda pointer */) ChangeName(to string) {
	c.Name = to
}
func (c *Cat /* pake '*' untuk penanda pointer */) ChangeAge(to int) {
	c.Age = to
}

func main() {
	var data Cat = Cat{
		Name: "Thad Deus Reginald",
		Age:  17,
	}
	fmt.Println(data)

	data.ChangeName("Marquis Evander Thornevale") // tidak perlu menggunakan '&' kalo di pointer in method
	fmt.Println(data)
}
