package models

import (
	"errors"
	"strings"
)

//Validate adalah fungsi untuk memberitahu bahwa inputan tidak boleh kosong
func Validate(input string) (bool, error) {
	//TrimSpace () digunakan untuk menghilangkan spasi untuk menanggulangi user hanya memberikan spasi
	if strings.TrimSpace(input) == "" {
		return false, errors.New("field tidak boleh kosong")
	}

	return true, nil
}
