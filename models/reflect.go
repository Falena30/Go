package models

import (
	"fmt"
	"reflect"
)

//GetProperyInfo adalah method untuk mengtahui info dari suatu variabel dengan menggunakan reflect
func (M *Member) GetProperyInfo() {
	var ReflectValue = reflect.ValueOf(M)

	if ReflectValue.Kind() == reflect.Ptr {
		ReflectValue = ReflectValue.Elem()
	}

	var ReflectType = ReflectValue.Type()

	for i := 0; i < ReflectValue.NumField(); i++ {
		//menyetak nama dari strcut member
		fmt.Println("Nama : ", ReflectType.Field(i).Name)
		//menyetak tipe data
		fmt.Println("Tipe Data : ", ReflectType.Field(i).Type)
		//menyetak nilai dari struck member
		fmt.Println("Umur : ", ReflectValue.Field(i).Interface())
		fmt.Println("")
	}
}
