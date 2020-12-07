package main

import (
	"fmt"

	"github.com/Go-learn/Web/Main/serve"
)

func main() {
	var nUser, err = serve.FetchUser()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	for _, each := range nUser {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}
}
