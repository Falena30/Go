package main

import (
	"encoding/json"
	"fmt"
)

//ini adalah path untk lokasi folder yang berbeda terlalu panjang, tapi bagaimana lagi
//var path = "/Users/User1/go/src/github.com/Go-learn/Web/templete/templete.html"

type user struct {
	//ingat penulisan json itu kecil
	FullName string `json:"Name"`
	Age      int
}

func main() {
	//ini adalah halam default dari webnya
	//handlefunc penentuan aksi jika user mengetik url tertentu dalam hal ini "/"
	//memiliki dua paramater link tujuan contoh : "/", dan callback atau aksi yang akan dilakukan jika url di akses
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			var data = map[string]string{
				//ingat harus sama dengan yang ditemplete
				"Name":     "Dimas Adi Suyikno",
				"Messsage": "Hi Selamat Berjuang belajar GO",
			}
			//jika level lokasi sama tinggal panggi nama dan extension dari filenya
			//untuk lebih mudah folder temple diletakkan dibawah folder main
			//ParseFiles, mengembalikan dua parameter salah satunya adalah error
			var t, err = template.ParseFiles("templete/templete.html")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			//Excute menjalankan perintah dan memiliki dua parameter, responswriter dan data
			t.Execute(w, data)
		})

		//jika menmabahkan index akan muncul func index
		http.HandleFunc("/index", server.Index)

		fmt.Println("Starting Web Server at http://localhost:8080/")
		//untuk menghiupkan server, memiliki dua parameter
		http.ListenAndServe(":8080", nil)
	*/

	//URL Parsing
	//merubah URL menjadi string
	/*
		var urlString = "http://jomblo.com:80/hello?name=dimas adi&age=24"
		var uParseURLString, err = url.Parse(urlString)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("url : %s\n", urlString)

		fmt.Printf("protocol : %s\n", uParseURLString.Scheme) // Protocol = http, prtotocol yang digunakan adalah http
		fmt.Printf("host : %s\n", uParseURLString.Host)       // jomblo.com:80, host dimana lokasi host dari server
		fmt.Printf("path : %s\n", uParseURLString.Path)       //hallo, atau path dari url sekarang

		var nama = uParseURLString.Query()["name"][0] //dimas adi
		var age = uParseURLString.Query()["age"][0]   //24
		fmt.Printf("Nama : %s, Umur : %s\n", nama, age)
	*/

	//JSON
	//Penggunaan JSON
	//deklarasi JSON, ada diatas
	var JSONString = `{
		"Name" : "Dimas Adi", 
		"Age" : 24
	}`

	var JSONData = []byte(JSONString)

	var data user
	//fungsi Unmarshal hayna bisa menerima data berbentu byte, dan pointer
	var err = json.Unmarshal(JSONData, &data)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Println("user : ", data.FullName)
	fmt.Println("age  : ", data.Age)

	var data1 map[string]interface{}
	//ingat salah satu parameter unmarshal ada pointer
	json.Unmarshal(JSONData, &data1)

	var data2 interface{}
	json.Unmarshal(JSONData, &data2)
	var decodeData2 = data2.(map[string]interface{})
	fmt.Println("user dengan interfce{}: ", decodeData2["Name"])
	fmt.Println("age dengan interfce{}: ", decodeData2["Age"])

	var ajsonString = `[
		{"Name": "john wick", "Age": 27},
    	{"Name": "ethan hunt", "Age": 32}
	]`

	var aData []user

	var aErr = json.Unmarshal([]byte(ajsonString), &aData)
	if aErr != nil {
		fmt.Println(aErr.Error())
		return
	}

	fmt.Println("user 1:", aData[0].FullName)
	fmt.Println("user 2:", aData[1].FullName)

	//Encode ke JSON
	var objectEncode = []user{{"john wick", 27}, {"ethan hunt", 32}}

	//marshal memiliki satu parameter
	var dJson, dErr = json.Marshal(objectEncode)

	if dErr != nil {
		fmt.Println(dErr.Error())
		return
	}
	var decodeString = string(dJson)
	fmt.Println(decodeString)
}
