package golanggoroutines

/**
Membuat Goroutine

● Untuk membuat goroutine di Golang sangatlah sederhana
● Kita hanya cukup menambahkan perintah go sebelum memanggil function yang akan kita jalankan
dalam goroutine
● Saat sebuah function kita jalankan dalam goroutine, function tersebut akan berjalan secara
asynchronous, artinya tidak akan ditunggu sampai function tersebut selesai
● Aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yang kita
buat selesai

//tidak cocok dijalankan pada func yang punya return value, karna return value tersebut tidak bisa ditangkap
*/

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

/**
Goroutine Sangat Ringan

● Seperti yang sebelumnya dijelaskan, bahwa goroutine itu sangat ringan
● Kita bisa membuat ribuan, bahkan sampai jutaan goroutine tanpa takut boros memory
● Tidak seperti thread yang ukurannya berat, goroutine sangatlah ringan
*/

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 1; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
