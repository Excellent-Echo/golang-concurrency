package concurrency

import "fmt"

func producer(c chan int) {
	for i := 1; i < 10; i++ {
		c <- i
	}
	close(c)

	// c <- 20
}

func closeRangeChan() {
	ch := make(chan int)

	go producer(ch) // [1,2,3,4,5,6,7,8,9,10,11]

	// for {
	// 	value, check := <-ch // check = false ketika close, true ketika masih diisi

	// 	if check == false {
	// 		fmt.Println("receive", value, check)
	// 		break
	// 	}
	// 	fmt.Println("receive", value, check)
	// }

	for value := range ch {
		fmt.Println("receive", value)
		// time.Sleep(500 * time.Millisecond) //scheduling
	}
}
