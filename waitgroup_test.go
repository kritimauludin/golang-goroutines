package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)
	fmt.Println("Hello")

	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go RunAsynchronous(wg)
	}

	wg.Wait()
	fmt.Println("Done")
}
