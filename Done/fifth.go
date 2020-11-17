package main

import "fmt"

var member = map[int]string{
	1234567: "Dimas",
	3715710: "Adi",
	8712074: "Suyikno",
}

func main() {
	//key value dengan map
	//map[type data untuk key] valuenya apa

	//kenapa jika melalkukan println sebeluum melakukan loop akan membuat println didalam loop akan terbalik?
	deleteID(1234567)
	fmt.Println(member)
}

func checkID(id int) bool {
	_, exist := member[id]

	return exist
}

func deleteID(id int) string {
	pesan := ""
	for idMember, _ := range member {
		if id == idMember {
			delete(member, id)
			pesan = "berhasil di hapus"
		} else {
			pesan = "Id Tidak ada"
		}
	}

	return pesan
}
