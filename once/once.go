package main

import (
	"fmt"
	"sync"
)

func main() {
	var x sync.Once

	for i := 0; i < 10; i++ {
		x.Do(func() {
			fmt.Println("Hello")
		})
	}
}
