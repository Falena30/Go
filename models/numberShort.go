package models

import "fmt"

//BubbleShort adalah sebuah algoritma sorting sederhana
func BubbleShort(angka []int) []int {

	for i := 0; i < len(angka)-1; i++ {
		fmt.Println(angka)
		for j := 0; j < len(angka)-1; j++ {
			if angka[j] > angka[j+1] {
				tampung := angka[j]
				angka[j] = angka[j+1]
				angka[j+1] = tampung
			}

		}
	}
	return angka
}

//InsertionShort adalah algoritma sorting
func InsertionShort(angka []int) []int {

	for i := 1; i < len(angka); i++ {
		ValueToInsert := angka[i] // menampung index dari array
		HolePosition := i         // position
		fmt.Println(angka)
		for HolePosition > 0 && angka[HolePosition-1] > ValueToInsert {
			angka[HolePosition] = angka[HolePosition-1]
			HolePosition = HolePosition - 1
		}
		angka[HolePosition] = ValueToInsert
	}
	return angka
}

//SelectionShort adalah algoritma sorting dengan cara memilih paling kecil dan langsung ditukar
func SelectionShort(angka []int) []int {

	for i := 0; i < len(angka)-1; i++ {
		min := i
		fmt.Println(angka)
		for j := i + 1; j < len(angka); j++ {
			if angka[j] < angka[min] {
				min = j
			}
		}
		if min != i {
			tampung := angka[min]
			angka[min] = angka[i]
			angka[i] = tampung
		}
	}
	return angka
}

//ShellShort shor merupakan salah satu type sort algorithm
func ShellShort(angka []int) []int {
	interval := 0

	for interval < len(angka)/3 {
		interval = interval*3 + 1
	}

	for interval > 0 {

		for outer := interval; outer < len(angka); outer++ {
			ValueToInsert := angka[outer]
			inner := outer
			for inner > interval-1 && angka[inner-interval] >= ValueToInsert {
				angka[inner] = angka[inner-interval]
				inner = inner - interval

			}
			angka[inner] = ValueToInsert
		}
		interval = (interval - 1) / 3

	}
	return angka
}
