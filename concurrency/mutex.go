package concurrency

import (
	"fmt"
	"sync"
)

var x int //0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x += 1
	m.Unlock()
	wg.Done()
}

func mutex() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, &m)
	}
	wg.Wait()

	// time.Sleep(1 * time.Second)

	fmt.Println("final total x", x)
}
