package main

import "fmt"

func main() {
	//array
	nama := [...]string{"Dimas", "Adi", "Suyikno", "Ahmad", "Aulia", "Wahib"}

	// catatan khusus range selalu return dua variabel, jadi jika ingin hanya satu yang diisi
	// harus mengganti salah satunya dengan "_"
	for index, value := range nama {
		// printf digunakan untuk menceak %d dan %s misalkan
		fmt.Printf("index ke- %d dan isinya adalah %s\n", index, value)
	}

}
