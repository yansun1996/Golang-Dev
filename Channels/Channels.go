package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // signal only channel

func main() {
	ch := make(chan int, 50) // internal data store 50 integers
	wg.Add(2)
	go func(ch <-chan int) {
		//for i := range ch {
		//	fmt.Println(i)
		//}
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 77
		ch <- 66
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()

	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
}

func logger() {
	//for entry := range logCh {
	//	fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:00:00"), entry.severity, entry.message)
	//}
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:00:00"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}
