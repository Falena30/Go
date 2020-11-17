package main

import (
	"fmt"
	"Go/models"
)

func main() {
	user1 := models.member("dimas",24)
	user1.umur = 23
	user2 := user1
	user2.nama = "Wahib"
	fmt.Println(user2)
}
