package main

import "fmt"

func main() {
	//key value dengan map
	//map[type data untuk key] valuenya apa
	var member = map[int]string{
		1234567: "Dimas",
		3715710: "Adi",
		8712074: "Suyikno",
	}

	for id, value := range member {
		fmt.Println(id, value)
	}
}
