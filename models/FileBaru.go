package models

import (
	"fmt"
	"io"
	"os"
)

//disini sekalian memberi nama dan tipenya
var path = "/Users/User1/go/src/github.com/Go-learn/main/text.txt"

//IsError befungsi untuk melakukan pengecekan
func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

//CreateFile adalah fungsi yang berfungsi untuk membuat file
func CreateFile() {
	//cek apakah file sudah ada atau belum
	var _, err = os.Stat(path)

	//buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if IsError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("File Telah Berhasil dibuat !", path)

}

//WriteFile berfungsi untuk mengisi file yang sebelumny telah dibuat
func WriteFile() {
	//buka file dengan akses write and read
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if IsError(err) {
		return
	}
	defer file.Close()
	//tulisa data yang ingin dimasukknan
	_, err = file.WriteString("halo\n")
	if IsError(err) {
		return
	}
	_, err = file.WriteString("ayo belajar Golang dari dasar\n")
	if IsError(err) {
		return
	}
	//simpan perubahan dengan sync
	err = file.Sync()
	if IsError(err) {
		return
	}
	fmt.Println("Yes file berhasil di rubah")
}

//ReadFile adalah fungsi yang digunkan untuk membaca suatu file
func ReadFile() {
	//buka dulu filenya
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)

	if IsError(err) {
		return
	}
	defer file.Close()

	//baca filenya
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if IsError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if IsError(err) {
		return
	}
	fmt.Println("=> file berhasil dibaca")
	fmt.Println(string(text))
}

//RemoveFile adalah fungsi untuk menghapus file
func RemoveFile() {
	var err = os.Remove(path)
	if IsError(err) {
		return
	}
	fmt.Println("File berhasil di hapus!")
}
