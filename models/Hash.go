package models

import (
	"crypto/sha1"
	"fmt"
	"time"
)

//DoHashUsingSalt adalah fungsi untuk mengdapi serangan karena hash bisa jadi memiliki nilai yang sama
//sangat berfungsi untuk enkripsi password user
func DoHashUsingSalt(text string) (string, string) {
	//memberitahukan waktu sekarang
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	//memberitahukan text dan saltnya
	var saltText = fmt.Sprintf("Text : '%s', salt : %s", text, salt)
	fmt.Println(saltText)
	//membuat variabel hash
	//strep 1
	var sha = sha1.New()
	//step 2 memberikan huruf acak
	sha.Write([]byte(saltText))
	var encrypt = sha.Sum(nil)
	return fmt.Sprintf("%x", encrypt), salt
}
