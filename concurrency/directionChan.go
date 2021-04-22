package concurrency

import "fmt"

func sender(c chan<- string) {
	c <- "ping"
	// test := <-c
}

func receiver(c <-chan string) {
	// c <- "pong"
	msg := <-c
	fmt.Println(msg)
}

func directionChan() {
	// channel direction

	// mengatur channel bisa menerima, mengirim, atau dua2nya
	ch := make(chan string)

	// ch <- "pong"
	// result := <-ch
	go sender(ch)
	go receiver(ch)

	// time.Sleep(1 * time.Second)
	// fmt.Println(result)
}
