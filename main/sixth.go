package main

import (
	"github.com/Go-learn/models"
)

func main() {

	user1 := models.Member{
		Nama: "dimas",
		Umur: 24,
	}
	//jadi jika variabel sudah di assignment maka variabel itu sudah memiliki perintahnya sendiri
	//jadi tidak perlu lagi memanggil packagenya
	user1.GetInfo()

	//fmt.Println(user1)
}
