package golanggoroutines

/**
Once

● Once adalah fitur di Go-Lang yang bisa kita gunakan untuk memastikan bahsa sebuah function di
eksekusi hanya sekali
● Jadi berapa banyak pun goroutine yang mengakses, bisa dipastikan bahwa goroutine yang pertama
yang bisa mengeksekusi function nya
● Goroutine yang lain akan di hiraukan, artinya function tidak akan dieksekusi lagi
*/
import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	wg := sync.WaitGroup{}
	once := sync.Once{}
	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1)
			once.Do(OnlyOnce)
			// OnlyOnce()
			wg.Done()

		}()
	}

	wg.Wait()
	fmt.Println("counter", counter)
}
