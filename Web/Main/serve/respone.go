package serve

import (
	"encoding/json"
	"net/http"

	"github.com/Go-learn/Web/Main/data"
)

var baseURL = "http://localhost:8080"

//Users berfungsi untuk meberikan response berupa menampilkan semua data yang ada
func Users(w http.ResponseWriter, r *http.Request) {
	//jenis tipe responnya adaalah json
	w.Header().Set("Content-Type", "appliaction/json")

	if r.Method == "POST" {

		var result, err = json.Marshal(data.DData)
		if err != nil {
			//http.Error memiliki 3 parameter yaitu response, error dan int
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//mendaftarkan data sebagai respone
		w.Write(result)
	}
	http.Error(w, "", http.StatusBadRequest)
}

//User befungsi untuk menampilkan data yang dicari
func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data.DData {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}

		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

//FetchUser digunakan untuk melakukan request ke localhost:8080, menerima response dari request tersebut, lalu menampikannya
func FetchUser() ([]data.Student, error) {
	var err error
	var client = &http.Client{}
	var data []data.Student
	requst, err := http.NewRequest("POST", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	respone, err := client.Do(requst)
	if err != nil {
		return nil, err
	}
	defer respone.Body.Close()

	err = json.NewDecoder(respone.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
