package models

import (
	"fmt"
	"strings"
)

//StringBubbleShort adalah sebuah algoritma sorting sederhana
func StringBubbleShort(Huruf []string) []string {

	for i := 0; i < len(Huruf)-1; i++ {
		fmt.Println(Huruf)
		for j := 0; j < len(Huruf)-1; j++ {
			if strings.Compare(Huruf[j], Huruf[j+1]) > 0 {
				tampung := Huruf[j]
				Huruf[j] = Huruf[j+1]
				Huruf[j+1] = tampung
			}

		}
	}
	return Huruf
}
