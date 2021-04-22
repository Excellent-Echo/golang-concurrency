package concurrency

import (
	"fmt"
	"time"
)

func server1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "from server 1"
}

func server2(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "from server 2"
}

func selectChannel() {
	// select
	output1 := make(chan string)
	// output2 := make(chan string)

	go server1(output1)
	// go server2(output2)

	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case s1 := <-output1:
			fmt.Println(s1)
			return
		default:
			fmt.Println("proses masih berjalan")
		}
	}

	// select {
	// case s1 := <-output1:
	// 	fmt.Println(s1)
	// case s2 := <-output2:
	// 	fmt.Println(s2)
	// default:
	// 	fmt.Println("proses masih berjalan")
	// }
}
