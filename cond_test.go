package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
*
Cond

● Cond adalah adalah implementasi locking berbasis kondisi.
● Cond membutuhkan Locker (bisa menggunakan Mutex atau RWMutex) untuk implementasi
locking nya, namun berbeda dengan Locker biasanya, di Cond terdapat function Wait() untuk
menunggu apakah perlu menunggu atau tidak
● Function Signal() bisa digunakan untuk memberi tahu sebuah goroutine agar tidak perlu menunggu
lagi, sedangkan function Broadcast() digunakan untuk memberi tahu semua goroutine agar tidak
perlu menunggu lagi
● Untuk membuat Cond, kita bisa menggunakan function sync.NewCond(Locker)
*/

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var wg = &sync.WaitGroup{}

func WaitCondition(value int) {
	defer wg.Done()
	wg.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	//secara lgsg mengirim signal ke seluruh antrian goroutine
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	wg.Wait()
}
