package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/Go-learn/models"
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

	var input string
	fmt.Print("masukkan namamu : ")
	fmt.Scanln(&input)

	if valid, err := models.Validate(input); valid {
		fmt.Println("Hallo : ", input)
	} else {
		fmt.Println(err.Error())
	}

}
