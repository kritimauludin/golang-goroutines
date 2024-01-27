package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
*
Once

● Once adalah fitur di Go-Lang yang bisa kita gunakan untuk memastikan bahsa sebuah function di
eksekusi hanya sekali
● Jadi berapa banyak pun goroutine yang mengakses, bisa dipastikan bahwa goroutine yang pertama
yang bisa mengeksekusi function nya
● Goroutine yang lain akan di hiraukan, artinya function tidak akan dieksekusi lagi
*/
func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("kriti")
	pool.Put("mauludin")
	pool.Put("aul")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("done")
}
