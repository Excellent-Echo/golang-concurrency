package concurrency

import (
	"fmt"
	"time"
)

// func sum(data []int, c chan int) {
// 	var sum int

// 	for _, val := range data {
// 		sum += val
// 	}
// 	c <- sum // sending data to channel
// }

func sum(data []int, c chan int) {

	var sum int
	for _, val := range data {
		sum += val
	}
	c <- sum // sending
}

func receive(c chan int) {
	result := <-c
	fmt.Println(result)
}

func channel() {
	// channel
	var v = []int{10, 20, 30, 10, 30, 20, 10} // 130
	var channel = make(chan int)
	go sum(v, channel)      // 130
	go receive(channel)     // 130
	go sum(v[0:2], channel) // 30
	go receive(channel)     // 30

	time.Sleep(1 * time.Second)

	fmt.Println("finish")

	// // declare channel
	// var c = make(chan int)

	// go sum(v[0:3], c) // 60
	// go sum(v[1:5], c) // 90

	// x := <-c
	// y := <-c

	// fmt.Println(x, y)
}
