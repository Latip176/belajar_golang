package main

import "fmt"

type Information struct {
	Data map[string]interface{}
}

func (info Information) cekData() interface{} {
	if info.Data["Name"] == "Latif Harkat" {
		return "Welcome!"
	} else {
		return nil
	}
}

func main() {
	data := map[string]interface{}{
		"Name": "Latif",
		"Age":  18,
	}
	info := Information{
		Data: data,
	}
	fmt.Println(info.cekData())

	// bisa di pake di kondisi
	if info.cekData() == nil {
		fmt.Println("Data tidak Ditemukan!")
	}
}
