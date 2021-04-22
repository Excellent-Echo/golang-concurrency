package concurrency

import (
	"fmt"
	"time"
)

func process(i int) {
	fmt.Println("started go routine", i)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Goroutine %d ended\n", i)
}

func waitGroup() {
	// var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		// wg.Add(1)
		go process(i)
	}
	// wg.Wait()
	time.Sleep(3 * time.Second)
	fmt.Println("all goroutine finished")
}
