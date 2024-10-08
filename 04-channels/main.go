package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			if i == 3 {
				message <- "three"
			}
		}
	}()

	msg := <-message
	fmt.Println(msg)
}
