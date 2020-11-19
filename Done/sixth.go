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

	kotak1 := models.Kotak{
		Panjang: 10,
		Lebar:   5,
	}

	models.HitunganBangunan(kotak1)

	lingakan := models.Lingkaran{
		Radius: 8,
	}

	models.HitunganBangunan(lingakan)
	//jadi jika variabel sudah di assignment maka variabel itu sudah memiliki perintahnya sendiri
	//jadi tidak perlu lagi memanggil packagenya
	//user1.GetInfo()
	user1.ChangeName("Suyikno", 22)
	fmt.Println(user1)
	//fmt.Println(user1)
}
