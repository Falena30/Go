package main

import "fmt"

func getPoint(point *int) {
	*point = 200
}
func main() {
	var point = 100
	getPoint(&point)

	fmt.Println(point)
}
