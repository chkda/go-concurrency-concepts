package main

import (
	"fmt"
	"sync"
	"time"
)

func HelloWorld(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 5)
	fmt.Println("Hello World")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go HelloWorld(&wg)
	go func() {
		fmt.Println("Hello Go")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("Hello Go-2")
		wg.Done()
	}()

	wg.Wait()
}
