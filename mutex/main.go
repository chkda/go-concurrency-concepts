package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	myMutex := sync.Mutex{}
	myMutex.Lock()
	go func() {
		myMutex.Lock()
		fmt.Println("I am in the go routine")
		myMutex.Unlock()
	}()
	fmt.Println("I am in the main routine")
	myMutex.Unlock()
	time.Sleep(time.Second * 1)
}
