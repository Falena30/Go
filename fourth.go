package main

import "fmt"

func main() {
	//array satu dimensi
	nama := []string{"Dimas", "Adi", "Suyikno", "Ahmad", "Aulia", "wahib"}
	// array multi dimensi
	//using make in slice atau insialisai slice tanpa mengambil dari array
	// catatan length dari slice harus sama atau setidaknya lebih besar dari array yang mau di copy
	newNama := make([]string, 6)
	//menggunakan copy dan append
	copy(newNama, nama)
	newNama[1] = "ADY"
	newNama = append(newNama, "friyadi", "udisnu")
	fmt.Println(newNama)
	fmt.Println(nama)
	//penggunaan slice [awal : ahiran]
	//jika menggunakan slice dan merubah isi dari variabel, variabel akan berubah juga
	//sliceNama := nama[2:4]

	// catatan khusus range selalu return dua variabel, jadi jika ingin hanya satu yang diisi
	// harus mengganti salah satunya dengan "_"

}
