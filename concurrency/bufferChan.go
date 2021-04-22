package concurrency

import "fmt"

func bufferChan() {
	// buffer channel
	// kita membatasi send / receive dari si channel

	// var c = make(chan int, 3)
	var c = make(chan int, 4)

	c <- 10 // [x , -, - , -]
	c <- 20 // [x,  x , -, -]
	c <- 30 // [x, x, x , -]
	c <- 40 // [x, x, x, x]
	// data1 := <-c // [- , x , x]
	// data2 := <-c // [- , - , x]
	// data3 := <-c // [-, - , - ]

	// c <- 50      // [x, x, -]
	// c <- 60      // [x, x, x]

	// fmt.Println(data1)
	// fmt.Println(data2)
	// fmt.Println(data3)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
