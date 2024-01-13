package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
}

var latifharkat = []Data{
	{
		ID:      "127.0.0.1",
		Name:    "Latif Harkat",
		Address: "Sukabumi, Jawa Barat, Indonesia",
		Age:     18,
	},
	{
		ID:      "127.0.0.1",
		Name:    "Latif Harkat",
		Address: "Sukabumi, Jawa Barat, Indonesia",
		Age:     18,
	},
	{
		ID:      "127.0.0.1",
		Name:    "Latif Harkat",
		Address: "Sukabumi, Jawa Barat, Indonesia",
		Age:     18,
	},
}

func getData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, latifharkat)
}

func main() {
	router := gin.Default()

	router.GET("/data", getData)
	router.Run("localhost:8080")
}
