package main

import (
	"fmt"

	"github.com/Go-learn/models"
)

func main() {
	user1 := models.Member{
		Nama: "dimas",
		Umur: 24,
	}

	fmt.Println(user1)
}
