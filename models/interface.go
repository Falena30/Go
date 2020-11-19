package models

import (
	"fmt"
	"math"
)

//1. buatlah sebuah interface
//2. buatlah sebuah struck
//3. buatlah sebuang fungsi
//4. buatlah sebuah method

//Bentuk adalah interface untuk luas dan keililing bangunan
type Bentuk interface {
	Keliling() float64
	Luas() float64
}

//Kotak berisi panjang dan lebar
type Kotak struct {
	Panjang, Lebar float64
}

//Keliling adalah method yang berfungsi untuk menghitung keliling dari kotak
func (K Kotak) Keliling() float64 {
	return 2*K.Panjang + 2*K.Lebar
}

//Luas adalah method yang berfungsi untuk menghitung luas dari kotak
func (K Kotak) Luas() float64 {
	return K.Panjang * K.Lebar
}

//Lingkaran berisi Radius
type Lingkaran struct {
	Radius float64
}

//Keliling adalah method yang berfungsi untuk menghitung keliling dari lingkaran
func (L Lingkaran) Keliling() float64 {
	return 2 * math.Pi * L.Radius
}

//Luas adalah method yang berfungsi untuk menghitung luas dari lingkaran
func (L Lingkaran) Luas() float64 {
	return math.Pi * L.Radius * L.Radius
}

// HitunganBangunan adalah fungsi untuk mencetak keliling dan luas dari interface
func HitunganBangunan(B Bentuk) {
	fmt.Println("luas dari banguan adalah ", B.Luas())
	fmt.Println("Keliling dari bangun adalah ", B.Keliling())
}
