package models

import "fmt"

//Member package berfungsi sebagai pointer
type Member struct {
	Nama string
	Umur int
}

//GetInfo method digunakan oleh user untuk mengetahui namanya
func (M Member) GetInfo() {
	fmt.Println("halo nama saya adalah")
}

//ChangeName merubah nama
func (M *Member) ChangeName(Name string, UmurBerubah int) {
	M.Nama = Name
	M.Umur = UmurBerubah
}
