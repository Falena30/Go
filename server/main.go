package main

import (
	"encoding/json"
	"fmt"
	"log"

	//"io/ioutil"
	"net/http"
)

type article struct {
	Tittle string `json :"judul"`
	Desc   string `json :"Deskribsi"`
}

//Articles adalah array dari article sturct
type Articles []article

var articleVariable = Articles{
	//karena ini termasuk struct, akhiran harus diberi koma
	article{Tittle: "Judul pertama", Desc: "ini adalah deskribisi dari judul pertama"},
	article{Tittle: "Judul keuda", Desc: "ini adalah deskribisi dari judul kedua"},
}

func main() {
	//ingat yang dijalankan adalah URLnya bukan methodnya
	http.HandleFunc("/", getHome)
	http.HandleFunc("/article", getArticles)
	http.HandleFunc("/post-articles", witLog(postArticles))
	http.ListenAndServe(":3000", nil)
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Testing respon pertama"))
}

func getArticles(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(articleVariable)
}
func postArticles(w http.ResponseWriter, r *http.Request) {
	//memberitahukan apakah request sama dengan method
	if r.Method == "POST" {
		//deklarasi variabel untuk menampung body
		//body, err := ioutil.ReadAll(r.Body)
		//if err != nil {
		//untuk memberitahu error dari internal server
		//	http.Error(w, "we cant read body", http.StatusInternalServerError)
		//}

		// response dari server dan mencetak bodynya
		//ingat untuk mengisi dari raw json yang dipakai adalah nama variabel bukan isinya
		var newArticles article
		err := json.NewDecoder(r.Body).Decode(&newArticles)

		if err != nil {
			fmt.Println("ada error yaitu", err)
		}

		//untuk memasukkan kedalam slice struck articlevariabel
		articleVariable = append(articleVariable, newArticles)

		//mencetak semua articleVariabel
		json.NewEncoder(w).Encode(articleVariable)
	} else {
		//mencetak kesalahan jika terjadie error
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func witLog(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//memberitahukan addres dari user
		log.Printf("Logged koneksi dari ", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
