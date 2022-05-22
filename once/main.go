package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func HelloWorld() {
	once.Do(func() {
		fmt.Println("Hello World - 1")
	})
	fmt.Println("Hello World")
}

func main() {
	HelloWorld()
	HelloWorld()
	HelloWorld()
}
