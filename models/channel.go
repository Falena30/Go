package models

import (
	"fmt"
	"math/rand"
	"time"
)

//PrintMassage adalah fungsi dengan parameter yang ada channelnya
func PrintMassage(what chan string) {
	fmt.Println(<-what)
}

//GetAverage adalah fungsi untuk menghitung rata2
func GetAverage(Angka []int, Ch chan float64) {
	var sum = 0

	for _, e := range Angka {
		sum += e
	}

	Ch <- float64(sum) / float64(len(Angka))
}

//GetMax merupakan fungsi untuk mengetahui nilai tertinggi
//channel direction bisa digunakan untuk mengirim dan menerima
func GetMax(Angka []int, Ch chan int) {
	var max = 0

	for _, e := range Angka {
		//melakukan kondisi apakan nilai dari e lebih besar dari max
		if max < e {
			max = e
		}
	}
	//memasukkan nilai max ke channel
	Ch <- max
}

//SendMassage adahal function channel yang mengirim pesan sebanyak 20 kali
//termasuk channel direction yaitu hanya bisa mengirim dalam kasusu ini mengirim ke CetakPesan()
func SendMassage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	//untuk menghentikan perulangan
	close(ch)
}

//CetakPesan adalah sebuah function channel dimana untuk mencetak pesan
//Channel direction berikut hanya bisa untuk menerima dalam kasus ini dari SendMessage()
func CetakPesan(ch <-chan string) {
	for pesan := range ch {
		fmt.Println(pesan)
	}
}

// SendData adalah function channel untuk mengirim data interval
func SendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		//jangan menggunakan crypto/ran tapi gunakan math/ran
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

//RetriveData adalah function untuk menerima data dari SendData()
func RetriveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print("Recive data ", data, " ", "\n")
		// case ini akan  terpenuhi jika terjadi interval 5 detik
		case <-time.After(time.Second * 5):
			fmt.Println("timeout. no activity for 5 second")
			break loop
		}
	}
}
