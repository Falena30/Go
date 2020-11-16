package main

import (
	"fmt"
)

//global variabel
var text = "text ini adalah global"

//fungsi
func multiplay(angka1 int, angka2 int) int {
	return angka1 * angka2
}

func getBio(umur int, nama string, status string) (biodataDiri string, umur10YAfter int) {
	biodataDiri = nama + " adalah seorang " + status
	umur10YAfter = umur + 10
	return biodataDiri, umur10YAfter
}

// function tanpa return
func printText() {
	fmt.Println(text)
}
func main() {
	fmt.Println("Nilai dari function adalah", multiplay(3, 33))
	basicInfo, AgeInfo := getBio(22, "Dimas", "jobless")
	fmt.Println(basicInfo)
	fmt.Println("umurnya setelah 10 tahun adalah ", AgeInfo)
	printText()
}
