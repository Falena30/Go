package main

import (
	"github.com/Go-learn/models"
)

type orang struct {
	Tinggi int
	Umur   int    `umur pada tahun ini`
	Nama   string `Nama dari orangnya`
}

// embed struct atau struct dalam struct
type murid struct {
	orang
	nilai    int
	kelas    int
	Predikat string
}

var angka = []int{4, 9, 14, 2, 6, 1}
var huruf = []string{"Dimas", "Dias", "Wahib", "Aulia", "Abdul"}

func main() {
	//fmt.Println(models.StringBubbleShort(huruf))
	/*AngkaCari := 22
	n := len(angka)
	result := models.BinarySearch(angka, 0, n-1, AngkaCari)

	if result == -1 {
		fmt.Println("angka tidak ada di array")
	} else {
		fmt.Println("Angka ada di array nome : ", result)
	}*/

	/*
		var s1 = murid{}
		s1.Nama = "Dimas Adi Suyikno"
		s1.Predikat = "Bagus"
		s1.Tinggi = 180
		s1.Umur = 17
		s1.kelas = 12
		s1.nilai = 80

		fmt.Println("Namaku : ", s1.orang.Nama, " Umur : ", s1.Umur, " Kelas : ", s1.kelas)
	*/

	var s1 = models.Kendaraan{
		NamaKendaraan:  "Civic",
		Roda:           4,
		TahunProduksi:  2017,
		JenisKendaraan: "Mobil",
	}

	s1.SayHello()

}
