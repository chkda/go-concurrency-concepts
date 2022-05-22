package main

import (
	"fmt"
	"sync"
)

func writeLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		for i := 0; i < 100; i++ {
			mux.Lock()
			m[i] = i
			fmt.Println("Write:", i)
			mux.Unlock()
		}
	}
}

func readLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		mux.RLock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mux.RUnlock()
	}
}

func main() {
	m := make(map[int]int)
	mutex := &sync.RWMutex{}
	go writeLoop(m, mutex)
	go readLoop(m, mutex)
	go readLoop(m, mutex)
	fmt.Scanln()
}
