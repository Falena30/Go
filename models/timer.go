package models

import (
	"fmt"
	"os"
	"time"
)

//Timer adalah fungsi untuk memberikan time out pada durasi yang diberikan
func Timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

//Watcher adalah fungsi untuk memberitahu bahwa waktu habis dan memaksa program untuk berhenti
func Watcher(timerout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out!, no answer more than ", timerout, " seconds")
	os.Exit(0)
}
