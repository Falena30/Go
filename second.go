package main

import (
	"fmt"
	"strconv"
)

//fungsi
func multiplay(angka1 int, angka2 int) int {
	return angka1 * angka2
}

func getBio(umur int, nama string, status string) (string, string) {
	umurTampung := strconv.Itoa(umur)
	return nama + " adalah seorang " + status,
		" yang berumur " + umurTampung + " tahun ini"
}
func main() {
	fmt.Println("Nilai dari function adalah", multiplay(3, 33))
	basicInfo, AgeInfo := getBio(22, "Dimas", "jobless")
	fmt.Println(basicInfo, AgeInfo)
}
