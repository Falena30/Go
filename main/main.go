package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// untuk membuat package menjadi selevel harus di kasih '.' sebelum ""

// package juga bisa dibuat menjadi alias seperti contoh
/*
	//jadi nanti kalau mau memanggil package models harus menulis mdl
	mdl "github.com/Go-learn/models"
*/

type orang struct {
	Tinggi int
	Umur   int    `umur pada tahun ini`
	Nama   string `Nama dari orangnya`
}

// embed struct atau struct dalam struct
type murid struct {
	orang
	nilai    int
	kelas    int
	Predikat string
}

type muridLengkap struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = muridLengkap{
	name:        "Dimas",
	height:      180.5,
	age:         24,
	isGraduated: true,
	hobbies:     []string{"Basketball", "Programming", "Gamming"},
}

var angka = []int{4, 9, 14, 2, 6, 1}
var huruf = []string{"Dimas", "Dias", "Wahib", "Aulia", "Abdul"}

func main() {
	//fmt.Println(models.StringBubbleShort(huruf))
	/*AngkaCari := 22
	n := len(angka)
	result := models.BinarySearch(angka, 0, n-1, AngkaCari)

	if result == -1 {
		fmt.Println("angka tidak ada di array")
	} else {
		fmt.Println("Angka ada di array nome : ", result)
	}*/

	/*
		var s1 = murid{}
		s1.Nama = "Dimas Adi Suyikno"
		s1.Predikat = "Bagus"
		s1.Tinggi = 180
		s1.Umur = 17
		s1.kelas = 12
		s1.nilai = 80

		fmt.Println("Namaku : ", s1.orang.Nama, " Umur : ", s1.Umur, " Kelas : ", s1.kelas)
	*/
	/*
		var s1 = models.Kendaraan{
			NamaKendaraan:  "Civic",
			Roda:           4,
			TahunProduksi:  2017,
			JenisKendaraan: "Mobil",
		}

		s1.SayHello()
	*/
	/*
		var s1 = &models.Member{"Dimas", 24}

		s1.GetProperyInfo()
	*/

	rand.Seed(time.Now().Unix())
	//pada saat dirunning menggunakan 2 proses
	runtime.GOMAXPROCS(2)

	//membuat channel dan di masukkan ke variable massage
	//var massage = make(chan string)

	//membuat buffer channel dan dimasukkan kedalam variable
	//massage := make(chan int, 2)

	/*
		//membaut closure yang dimasukkan ke varibale hello
		var Hello = func(who string) {
			//mencetak data dari who
			var data = fmt.Sprintf("hello %s", who)
			//deklarasi channel
			massage <- data
		}

		//Penggunaan Goroutine
		go Hello("Dimas Adi Suyikno")
		go Hello("Ahmad Aulia Wahib")
		go Hello("Imam Abdul Wahid")

		//melakukan channel
		var m1 = <-massage
		fmt.Println(m1)

		var m2 = <-massage
		fmt.Println(m2)

		var m3 = <-massage
		fmt.Println(m3)
	*/

	/*
		go func() {
			for {
				i := <-massage
				fmt.Println("recive data", i)
			}
		}()

		for i := 0; i < 5; i++ {
			fmt.Println("send data", i)
			massage <- i
		}
	*/

	/*
		for _, each := range []string{"Dimas", "Wahib", "Imam"} {
			go func(who string) {
				//mencetak data dari who
				var data = fmt.Sprintf("hello %s", who)
				//deklarasi channel
				massage <- data
			}(each)
		}

		for i := 0; i < 3; i++ {
			models.PrintMassage(massage)
		}
	*/

	//channel select
	/*
		var angka = []int{1, 3, 5, 6, 1, 7, 0, 12, 4, 56, 7}
		fmt.Println("angka : ", angka)

		var ch1 = make(chan float64)
		go models.GetAverage(angka, ch1)

		var ch2 = make(chan int)
		go models.GetMax(angka, ch2)

		for i := 0; i < 2; i++ {
			select {
			case avg := <-ch1:
				fmt.Printf("avg \t : %.2f \n", avg)
			case max := <-ch2:
				fmt.Printf("max \t : %d \n", max)
			}
		}
	*/

	//channel untuk range and close
	/*
		var massage = make(chan string)
		go models.SendMassage(massage)
		models.CetakPesan(massage)
	*/

	/*
		//channel untuk interval
		var massage = make(chan int)
		go models.SendData(massage)
		models.RetriveData(massage)
	*/

	//defer dan exit
	/*
		models.OrderMenu("Pizza")
		models.OrderMenu("Burger")
	*/

	//error default
	/*
		var input string
		fmt.Print("masukkan angaka : ")
		fmt.Scanln(&input)

		var number int
		var err error
		number, err = strconv.Atoi(input)

		if err == nil {
			fmt.Println(number, "is number")
		} else {
			fmt.Println(input, "isn't number")
			fmt.Println(err.Error())
		}*/
	/*
		defer models.Catach()
		var input string
		fmt.Print("masukkan namamu : ")
		fmt.Scanln(&input)

		if valid, err := models.Validate(input); valid {
			fmt.Println("Hallo : ", input)
		} else {
			panic(err.Error())
			fmt.Println("end")
		}
	*/

	/*
		//penggunan recover IIFE
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("terjadi error", r)
			} else {
				fmt.Println("aplikasi berjalan lancar")
			}
		}()

		panic("some error happen")
	*/

	/*
		list := []string{"Dimas", "Adi", "Suyikno", "Imam", "Abdul"}

		for _, each := range list {
			func() {
				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Panic occured on looping", each, "| message:", r)
					} else {
						fmt.Println("Application running perfectly")
					}
				}()
				panic("some error happen")
			}()

		}
	*/

	//mencetak macam - macam tipe data
	//penggunaan printf untuk menggunakannya
	//%b digunakan untuk mencetak biner dari data numeriknya : 01001
	fmt.Printf("Halo umur Saya : %b\n", data.age)
	//%c kode unicode di format menjadi data numerik contoh = 1400 -> n
	fmt.Printf("%c\n", 1400)
	//%d digunakan untuk memformat data numerik menjadi string numerik berbasis 10
	fmt.Printf("%d\n", data.age)
	//%e dan %E digunakan untuk memformat data numerik desimal menjadi bentuk numerik standard scientific notation atau yang 1.825000E+02
	fmt.Printf("%e\n", data.height)
	fmt.Printf("%E\n", data.height)
	//%f dan %f figunakan untuk memformat data numerik desimal, dengan lebar desimalnya bisa diatur
	fmt.Printf("%.f\n", data.height)  //bulat atau tidak ditampilkan bilangan decimalya
	fmt.Printf("%2.f\n", data.height) //menampilkan dua bilangan dibelakang ,
	fmt.Printf("%f\n", data.height)   //menampilkan bilangan decimal 6 digit
	fmt.Printf("%.9f\n", data.height) //menampilkan bilngan decimal 9 digit
	//%g atau %G sama seperti %f akan tetapi memliki lebar desimal yang lebih luas dan memiliki perbedaan lain yaitu lebar dari nilainya tidak bisa di custom
	fmt.Printf("%g\n", data.height)
	//%o merubah data numerik menjadi string numerik berbasis 8 (oktal) / bit
	fmt.Printf("%o\n", data.age)
	//%p menncetak pointer dari variabelnya, berbasis 16 dengan prefix 0x., ingat selalu menggunakan & sebelum nama variabel untuk pointer
	fmt.Printf("%p\n", &data.name)
	//%q digunakan untuk esacpe string, meskipun string dipakai menggugnakan literal '\' akan tetap di escape
	fmt.Printf("%q\n", `" name \ height "`)
	//%s digunakan untuk memformat data string
	fmt.Printf("%s\n", data.name)
	//%t digunakan untuk memformat data boolean, ditampilkan nilainya
	fmt.Printf("%t\n", data.isGraduated)
	//%T mengambil tipe variable yang akan diformat
	fmt.Printf("%T\n", data.name)
	fmt.Printf("%T\n", data.height)
	fmt.Printf("%T\n", data.age)
	fmt.Printf("%T\n", data.isGraduated)
	fmt.Printf("%T\n", data.hobbies)
	//%v digunakan untuk memformat semua tipe data bahkan inteface{}, jika struct yang dicetak maka akan menampilkan semuanya
	fmt.Printf("%v\n", data)
	//%+v digunakan untuk memformat struct, dan ditampilkan semuanya secara berurutan termasuk elemnya
	//khusus anonymous struct, maka akan muncul struktur anonymousnya yaitu strcut{isinya}
	fmt.Printf("%+v\n", data)
	//%x dan %X digunakan untuk memformat data numerik menjadi bentuk string numerik berbasis 16 atau hexadecimal
	fmt.Printf("%x\n", data.age)
	//%% untuk memformat % akan muncul pada saat di jalankan
}
