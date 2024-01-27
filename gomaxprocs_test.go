package golanggoroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

/**
GOMAXPROCS

● Sebelumnya diawal kita sudah bahas bahwa goroutine itu sebenarnya dijalankan di dalam Thread
● Pertanyaannya, seberapa banyak Thread yang ada di Go-Lang ketika aplikasi kita berjalan?
● Untuk mengetahui berapa jumlah Thread, kita bisa menggunakan GOMAXPROCS, yaitu sebuah
function di package runtime yang bisa kita gunakan untuk mengubah jumlah thread atau
mengambil jumlah thread
● Secara default, jumlah thread di Go-Lang itu sebanyak jumlah CPU di komputer kita.
● Kita juga bisa melihat berapa jumlah CPU kita dengan menggunakan function runtime.NumCpu()
*/

func TestGetDataGomaxprocs(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			wg.Done()
		}()
	}
	totalCpu := runtime.NumCPU()
	fmt.Println("Total cpu", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total goroutine", totalGoroutine) //minimal 2 yaitu untuk jalankan code dan garbage collectornya

	wg.Wait()
}

/*
*
Peringatan

● Menambah jumlah thread tidak berarti membuat aplikasi kita menjadi lebih cepat
● Karena pada saat yang sama, 1 CPU hanya akan menjalankan 1 goroutine dengan 1 thread
● Oleh karena ini, jika ingin menambah throughput aplikasi, disarankan lakukan vertical scaling
(dengan menambah jumlah CPU) atau horizontal scaling (menambah node baru)
*/
func TestChangeThread(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			wg.Done()
		}()
	}
	totalCpu := runtime.NumCPU()
	fmt.Println("Total cpu", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total goroutine", totalGoroutine) //minimal 2 yaitu untuk jalankan code dan garbage collectornya

	wg.Wait()
}
