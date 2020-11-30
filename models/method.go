package models

import "fmt"

//Kendaraan adalah struct yang berisi global kendaraan
type Kendaraan struct {
	Roda           int
	TahunProduksi  int
	NamaKendaraan  string
	JenisKendaraan string
}

//SayHello adalah method untuk mencetak struct kendaraan
func (K Kendaraan) SayHello() {
	fmt.Println("Nama Kendaraan : " + K.NamaKendaraan)
}
