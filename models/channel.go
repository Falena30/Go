package models

import "fmt"

//PrintMassage adalah fungsi dengan parameter yang ada channelnya
func PrintMassage(what chan string) {
	fmt.Println(<-what)
}
