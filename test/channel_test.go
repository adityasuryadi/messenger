package test

import (
	"fmt"
	"sync"
	"testing"
)

func print(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("message")
}

func TestChannel(t *testing.T) {
	// var ch = make(chan string)
	var wg sync.WaitGroup

	// go func() {
	for i := 0; i < 20; i++ {
		// ch <- fmt.Sprintf("hello %d", i)
		go print(&wg)
		wg.Add(1)
	}
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 20; i++ {
	// 		message := <-ch
	// 		fmt.Println(message)
	// 	}
	// }()
	// for m := range ch {
	// 	fmt.Println("ch", m)
	// }

	wg.Wait()
	fmt.Println("selesai")
}
