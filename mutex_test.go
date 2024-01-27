package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
*
Mutex (Mutual Exclusion)

● Untuk mengatasi masalah race condition tersebut, di Go-Lang terdapat sebuah struct bernama
sync.Mutex
● Mutex bisa digunakan untuk melakukan locking dan unlocking, dimana ketika kita melakukan
locking terhadap mutex, maka tidak ada yang bisa melakukan locking lagi sampai kita melakukan
unlock
● Dengan demikian, jika ada beberapa goroutine melakukan lock terhadap Mutex, maka hanya 1
goroutine yang diperbolehkan, setelah goroutine tersebut melakukan unlock, baru goroutine
selanjutnya diperbolehkan melakukan lock lagi
● Ini sangat cocok sebagai solusi ketika ada masalah race condition yang sebelumnya kita hadapi
*/
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

/*
*
RWMutex (Read Write Mutex)

● Kadang ada kasus dimana kita ingin melakukan locking tidak hanya pada proses mengubah data,
tapi juga membaca data
● Kita sebenarnya bisa menggunakan Mutex saja, namun masalahnya nanti akan rebutan antara
proses membaca dan mengubah
● Di Go-Lang telah disediakan struct RWMutex (Read Write Mutex) untuk menangani hal ini, dimana
Mutex jenis ini memiliki dua lock, lock untuk Read dan lock untuk Write
*/
type BankAccount struct {
	RWMutex sync.RWMutex
	balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.balance = account.balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.balance
	account.RWMutex.RUnlock()

	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println("last balance : ", account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("final balance : ", account.GetBalance())
}

/**
Deadlock

● Hati-hati saat membuat aplikasi yang parallel atau concurrent, masalah yang sering kita hadapi
adalah Deadlock
● Deadlock adalah keadaan dimana sebuah proses goroutine saling menunggu lock sehingga tidak
ada satupun goroutine yang bisa jalan
● Sekarang kita coba simulasikan proses deadlock
*/

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) ChangeBalance(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1 :", user1.Name)
	user1.ChangeBalance(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2 :", user2.Name)
	user2.ChangeBalance(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Kriti",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Mauludin",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 500000)

	time.Sleep(5 * time.Second)

	fmt.Println("User", user1.Name, "Balance", user1.Balance)
	fmt.Println("User", user2.Name, "Balance", user2.Balance)
}
