package main

import (
	"fmt"
	"strconv"
)

//import

//variable type here
var namaDepan, namaBelakang = "Dimas", "Adi Suyikno"

// tipe dari variabelnya
var angka int
var huruf string
var boolean bool

func main() {
	//shorthand untuk variabel
	//buju := "hallo"
	//angka = 10
	//angka2 := 17
	//angka3 := 4

	jumlahTotal := "300"
	penampungJumlahTotal, _ := strconv.Atoi(jumlahTotal)

	penambahan := penampungJumlahTotal + 77

	fmt.Println("jumlah akhir ", penambahan)
	//operasi matematika
	//var hasil int

	//hasil = angka * angka2 / angka3
	//fmt.Println("hasil dari perhitungan " + strconv.Itoa(hasil))

	//constan variabel
	//const hidup = 40

	//gunakan , untuk memisahkan int
	//fmt.Println("angka pertama", angka, " angka kedua ", angka2, " angka ketiga ", angka3)
	//gunakan + untuk memisahkan string
	//fmt.Println("Hallo saya " + namaDepan + " " + namaBelakang + " noob sedang belajar Golang")
}
