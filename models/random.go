package models

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//RandomString adalah fungsi untuk melakukan random string
func RandomString(lenght int) string {
	b := make([]rune, lenght)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
