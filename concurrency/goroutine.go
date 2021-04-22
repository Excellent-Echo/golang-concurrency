package concurrency

import (
	"fmt"
	"time"
)

func memindah() {
	fmt.Println("memindahkan barang ke gudang")
}

func menurunkan() {
	fmt.Println("menurunkan barang dari cargo")
}

func mendata() {
	fmt.Println("mendata barang di cargo")
}

func concurrent() {
	// cocurrency != parallel / async

	// goroutine // dasar terpenting concurrency go
	// 1. kita menampungnya ke dalam sebuah fungsi
	// 2. kita menggunakan sintaks "go"

	// go f("mendata")
	// go f("menurunkan")
	// go f("memindah")

	for i := 0; i < 10; i++ {
		go mendata()
		go menurunkan()
		go memindah()
	}

	// datanya ada 1 miliar data
	// scan link yg ada di website google (mengecek isi, mengecek tulisan, mengecek lainnya)

	//mendata() => menurunkan() => memindah()

	// 3 jalan bersamaan

	time.Sleep(1000 * time.Millisecond) // 1s = 1000 ms

	fmt.Println("finish")
}
