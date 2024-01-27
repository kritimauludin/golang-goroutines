package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
Timer

● Timer adalah representasi satu kejadian
● Ketika waktu timer sudah expire, maka event akan dikirim ke dalam channel
● Untuk membuat Timer kita bisa menggunakan time.NewTimer(duration)
*/

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

/*
*
time.After()

● Kadang kita hanya butuh channel nya saja, tidak membutuhkan data Timer nya
● Untuk melakukan hal itu kita bisa menggunakan function time.After(duration)
*/
func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

/**
time.AfterFunc()

● Kadang ada kebutuhan kita ingin menjalankan sebuah function dengan delay waktu tertentu
● Kita bisa memanfaatkan Timer dengan menggunakan function time.AfterFunc()
● Kita tidak perlu lagi menggunakan channel nya, cukup kirim kan function yang akan dipanggil
ketika Timer mengirim kejadiannya*/

func TestAfterFunc(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		wg.Done()
	})

	fmt.Println(time.Now())
	wg.Wait()
}
