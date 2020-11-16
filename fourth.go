package main

import "fmt"

func main() {
	//array

	var mobil [3]string

	mobil[0] = "Honda"
	mobil[1] = "BMW"
	mobil[2] = "Nissan"

	angka := [...]int{1, 4, 51, 6}

	angka[2] = 8
	fmt.Println(mobil)
	fmt.Println(angka)
}
