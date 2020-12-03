package models

import "fmt"

//OrderMenu adalah fungsi untuk memesan makanan
func OrderMenu(Menu string) {
	defer fmt.Println("Silahkan tunggu makannanya!")

	if Menu == "Pizza" {
		fmt.Print("Pilihan tepat!", " ")
		fmt.Print("Pizza ditempat kami paling enak!", "\n")
		return
	}
	fmt.Println("pesanan anda : ", Menu)
}
