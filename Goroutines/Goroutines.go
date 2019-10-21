package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	// go sayHello()
	// time.Sleep(100 * time.Millisecond)

	// // msg is changed, result is changed
	// var msg = "Hello"
	// go func() {
	// 	fmt.Println(msg)
	// }()
	// msg = "Goodbye"
	// time.Sleep(100 * time.Millisecond)

	// // msg2 as parameter, result not changed
	// var msg2 = "Hello"
	// go func(msg string) {
	// 	fmt.Println(msg)
	// }(msg2)
	// msg2 = "Goodbye"
	// time.Sleep(100 * time.Millisecond)

	// // Use waitgroup to synchronize
	// var msg3 = "Hello"
	// wg.Add(1)
	// go func(msg string) {
	// 	fmt.Println(msg)
	// }(msg3)
	// msg3 = "Goodbye"
	// wg.Wait()

	// Multiple goroutine
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()

	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
