package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

/**
time.Ticker

● Ticker adalah representasi kejadian yang berulang
● Ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel
● Untuk membuat ticker, kita bisa menggunakan time.NewTicker(duration)
● Untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop()
*/

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	x := 1

	for time := range ticker.C {
		fmt.Println(time)
		if x <= 5 {
			x += 1
		} else {
			ticker.Stop()
		}

	}
}

/*
*
time.Tick()

● Kadang kita tidak butuh data Ticker nya, kita hanya butuh channel nya saja
● Jika demikian, kita bisa menggunakan function timer.Tick(duration), function ini tidak akan
mengembalikan Ticker, hanya mengembalikan channel timer nya saja
*/
func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}
