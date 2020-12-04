package main

import (
	"fmt"
	"math/rand"
	"regexp"
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

	/*
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
	*/

	/*
		//random
		//isi dari seednya adalah unix nano dari waktu sekarang
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Println("berikut adalah angka random ke-1 : ", rand.Intn(9))
		fmt.Println("berikut adalah angka random ke-2 : ", rand.Intn(6))
		fmt.Println("berikut adalah angka random ke-3 : ", rand.Intn(8))
	*/
	/*
		fmt.Printf("%s\n", models.RandomString(10))
	*/
	/*
		//penggunakan time.Now() yang di masukkan ke varible waktuSekarang
		var now = time.Now()
		//fmt.Printf("Sekarang jam : %v\n", waktuSekarang)
		fmt.Println("year : ", now.Year(), " Bulan : ", now.Month(), " day : ", now.Day(), " Hari : ", now.Weekday(), " Lokasi : ", now.Location())
		//penggunaan waktu secara manual , time.Date(tahun, bulan, tanggal, jam, menit, detik, nanodetik, timezone)
		var waktuSekarang2 = time.Date(2020, 12, 4, 8, 33, 0, 0, time.UTC)
		fmt.Printf("Sekarang jam : %v\n", waktuSekarang2)
	*/
	/*
		//penggunan parsing pada time.Time
		var value, LayoutString string
		var date time.Time

		LayoutString = "2006-01-02 15:04:05"
		value = "2020-12-04 08:46:05"

		date, _ = time.Parse(LayoutString, value)

		fmt.Println(value, "\t->", date.String())

		//LayoutString = "02/01/2006 MST"
		//value = "04/12/2020 WIB"
		date, _ = time.Parse(time.RFC822, "04 Dec 20 08:00 WIB")
		//mengganti tipe atau layout dari waktu
		var date1 = date.Format("Monday 02, January 2006 15:04 MST")

		fmt.Println("date1 : ", date1)

		var date2 = date.Format(time.RFC3339)
		fmt.Println("date2 : ", date2)
		var tanggal, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")

		//untuk handle error dalam pengisian tanggal
		if err != nil {
			fmt.Println("error : ", err.Error())
			return
		}
		fmt.Println(tanggal)
	*/

	//penggunaan Timer, Ticker & Schedule
	//time.sleep adalah menghentikan sementara program selama beberapa detik
	/*
		fmt.Println("Start")
		time.Sleep(time.Second * 5)
		fmt.Println("print again after 5 second")
		nilai := 0
		benar := true
		for benar {
			nilai++
			fmt.Println("Hallo!")
			time.Sleep(1 * time.Second)
			if nilai > 5 {
				benar = false
			}
		}
	*/
	//fungsi NewTimer() mengembalikan nilai *time.Time dan mengirim channel bernama C
	/*
		var newTime = time.NewTimer(4 * time.Second)
		fmt.Println("start")
		<-newTime.C
		fmt.Println("Finish")
	*/
	//fungsi afterfunc() fungsi ini berguna untuk penggunaan sebelum fungsi dilakukan
	/*
		var chanel = make(chan bool)
		time.AfterFunc(4*time.Second, func() {
			fmt.Println("expired")
			chanel <- true
		})
		fmt.Println("start")
		<-chanel
		fmt.Println("finish")
	*/
	//After() adalah sebuah fungsi pada paket time yang berfungsi hampir sama seperti sleep, yang membedakannya adalah After() mengembalikan nilai channel jadi harus mengunakna '<-'
	/*
		fmt.Println("start")
		<-time.After(3 * time.Second)
		fmt.Println("finish")
	*/
	/*
		done := make(chan bool)
		ticker := time.NewTicker(time.Second)
		//fungsi anonymous ini akan berjalan sendiri karena goroutine sehingga
		//jika terjadi 10 detik fungsi ini akan berjalan sendiri dan menghentikan
		//proses print hallo
		go func() {
			time.Sleep(10 * time.Second) //berhenti sejenak 10 detik
			done <- true
		}()

		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case t := <-ticker.C:
				fmt.Println("Hallo!", t)
			}
		}
	*/
	//kombinasi timer dan goroutine
	/*
		var timeout = 5
		var ch = make(chan bool)
		go models.Timer(timeout, ch)
		go models.Watcher(timeout, ch)

		var input string
		fmt.Print("berapa 50/5 ? ")
		fmt.Scan(&input)

		if input == "10" {
			fmt.Println("congrates you answer is correct!")
		} else {
			fmt.Println("Upps your answer is incorrect!")
		}
	*/
	/*
		start := time.Now()
		time.Sleep(5 * time.Second)
		//durasinya menggunakan Since() dimana ini dimulai dari awal dalam paramernya dibuat
		duration := time.Since(start)

		fmt.Println("time elapsed in seconds : ", duration.Seconds())
		fmt.Println("time elapsed in minutes : ", duration.Minutes())
		fmt.Println("time elapsed in hours	 : ", duration.Hours())
	*/
	/*
		//kalkulasi perbedaan waktu pada dua object, sama sih hasilnya dengan yang diatas
		t1 := time.Now()
		time.Sleep(5 * time.Second)
		t2 := time.Now()
		duration := t2.Sub(t1)
		fmt.Println("time elapsed in seconds : ", duration.Seconds())
		fmt.Println("time elapsed in minutes : ", duration.Minutes())
		fmt.Println("time elapsed in hours	 : ", duration.Hours())
		//mebuat durasi menggunakan variabel
		var n time.Duration = 5
		duration = n * time.Second // 5 detik
		//atau bisa seperti ini sama seperti diatas
		z := 5
		duration = time.Duration(z) * time.Second
	*/
	/*
		//strconv.Atoi(), dari string ke int, dan mengembalikan dua parameter error dan hasil kembalian
		//strconv.Itoa(), dari int ke string
		//strconv.ParseInt()
		var str = "123"
		var num, err = strconv.ParseInt(str, 10, 64)
		if err == nil {
			fmt.Println(num)
		}
	*/
	/*
		//konversi dari string byte
		var text = "halo"
		//dimasukkan perkata menjadi byte
		var b = []byte(text)
		//mencetaknya satu per satu
		fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])
		//kebalikannya
		var bByte = []byte{104, 97, 108, 111}
		var s = string(bByte)
		fmt.Printf("%s\n", s)

		var c int64 = int64('h')
		fmt.Printf("%d\n", c)
		var d string = string(104)
		fmt.Printf("%s\n", d)
	*/
	//type assertion pada interface kosong
	/*
		var dData = map[string]interface{}{
			"nama":    "Dimas Adi",
			"grade":   2,
			"height":  180.5,
			"isMale":  true,
			"hobbies": []string{"Gaming", "Programming"},
		}
	*/
	/*
		fmt.Println(dData["nama"].(string))
		fmt.Println(dData["grade"].(int))
		fmt.Println(dData["height"].(float64))
		fmt.Println(dData["isMale"].(bool))
		fmt.Println(dData["hobbies"].([]string))
	*/
	//cara ini lebih mudah jika terdapat banyak isi dalam interface
	/*
		for _, val := range dData {
			switch val.(type) {
			case string:
				fmt.Println(val.(string))
			case int:
				fmt.Println(val.(int))
			case float64:
				fmt.Println(val.(float64))
			case bool:
				fmt.Println(val.(bool))
			case []string:
				fmt.Println(val.([]string))
			}
		}
	*/
	//apa aja yang ada di paket string
	/*
		//contain(), memiliki dua parameter, target, dan yang dicari
		var nama string = "Dimas Adi"
		var isExist = strings.Contains(nama, "Z")
		fmt.Println(isExist)
		//hasPrefix(), adalah pencarian apakah awalannya sesuai dengan yang dicari atau tidak
		//Besar kecil juga berpengaruh jika tidak menggunakan Uppercase
		var isPrefix = strings.HasPrefix(nama, "Dim")
		fmt.Println(isPrefix)
		//hasSuffix(), kebalikan dari hasPrefix yaitu apakah akhir dari kalimatnya memiliki kata yang dicari
		var isSuffix = strings.HasSuffix(nama, "dim")
		fmt.Println(isSuffix)
		//count(), untuk menghitung jumlah kata yang dicari pada target
		//mengembalikan dalam bentuk numerik
		var isCount = strings.Count(nama, "i")
		fmt.Println(isCount)
		//indrx(), untuk mencari ke index keberapa
		//besar kecil huruf juga berpengaruh
		var isIndex = strings.Index(nama, "d")
		fmt.Println(isIndex)
		//replace(), mengganti string yang ditarget, bisa satu kata dua kata atau bahkan semuanya, varibalenya (target, cari, pengganti, pengganti 1 2 atau keseluruahn)
		var find = "i"
		var rReplace = "u"
		//-1 disini untuk merubah semuanya
		var isReplace = strings.Replace(nama, find, rReplace, -1)
		fmt.Println(isReplace)
		//repeat(), untuk mengulang string, memiliki dua parametr, target dan banyaknya jumlah di ulangnya
		//split(), string untuk memecah suatu kalima atau kata
		var string1 = strings.Split("Dimas Adi Suyikno", " ")
		fmt.Println(string1)
		var string2 = strings.Split("Dimas Adi Suyikno", "")
		fmt.Println(string2)
		//join(), untuk menggabungkan string, yang sebelumnya berupa slice atau array menjadi satu parameter string saja
		var dJoin = []string{"Imam", "Abdul", "Wahid"}
		var sJoin = strings.Join(dJoin, "-")
		fmt.Println(sJoin)
		//ToLower(), digunakan untuk membuat string huruf kecil semua
		var sToLower = strings.ToLower(nama)
		fmt.Println(sToLower)
		//ToUpper(), digunakan untuk membuat string huruf besar semua
		var sToUpper = strings.ToUpper(nama)
		fmt.Println(sToUpper)
	*/
	//Regex
	//findallstring(), digunakan untuk membuat satu kalimat menjadi pecahan slice perkata
	var txt = "kucing kucang kucucung"
	//compile seluruh huruf a-z dan dimasukkan ke regex
	//jika ada huruf besar maka tidak akan muncul hurufnya dari hasil regexnya
	var regex, err = regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Println(err.Error())
	}
	var res1 = regex.FindAllString(txt, 2)
	fmt.Println(res1)

	var res2 = regex.FindAllString(txt, -1)
	fmt.Println(res2)
	var isMatch = regex.MatchString(txt)
	fmt.Println(isMatch)
	//findstring(), digunakan untuk menemukan string yang dicari, jika terdapat kesamaan maka yang awal saja yang digunakan
	var fFind = regex.FindString(txt)
	//kucing karena kata pertama
	fmt.Println(fFind)
	//finstringindex(), digunakan untuk mencari index dari string yang dicari dari regex
	var sFind = regex.FindStringIndex(txt)
	//akan mencetak [0;6] artinya index dari 0 ke 6 yaitu kucing
	fmt.Println(sFind)
	var str = txt[0:6]
	fmt.Println(str)
	//findallstring, method ini digunakan untuk mencari semua string yang cocok dengan regex
	var aFind = regex.FindAllString(txt, -1)
	fmt.Println(aFind)
	//replace all string, merubah semua string
	var aReplace = regex.ReplaceAllString(txt, "kuceng")
	fmt.Println(aReplace)
	//replaceallstringfunc, merubah semua string tetapi dapat memiliki kondisinya sendiri
	var rRepalce = regex.ReplaceAllStringFunc(txt, func(each string) string {
		if each == "kucing" {
			return "kuceng"
		}
		return each
	})

	fmt.Println(rRepalce)
	//split, digunakan untuk memisahkan
	var rRegex, _ = regexp.Compile(`[c-k]+`) //yang dicari adalah c-k sebagai pemisah
	var aSplit = rRegex.Split(txt, -1)
	fmt.Printf("%#v\n", aSplit)

}
