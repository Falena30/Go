package models

import "fmt"

//Kendaraan adalah struct yang berisi global kendaraan(bisa dipanggil karena public)
type Kendaraan struct {
	//untuk bisa dipanggil pada file lain harus berupa public atau huruf besar dari nama variablenya
	Roda           int
	TahunProduksi  int
	NamaKendaraan  string
	JenisKendaraan string
}

//SayHello adalah method untuk mencetak struct kendaraan
func (K Kendaraan) SayHello() {
	fmt.Println("Nama Kendaraan : " + K.NamaKendaraan)
}
