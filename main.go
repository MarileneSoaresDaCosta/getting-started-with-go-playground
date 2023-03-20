package main

import (
	"fmt"
	"github.com/MarileneSoaresDaCosta/getting-started-with-go-playground/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}
