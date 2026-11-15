package main

import (
	"fmt"
	"sync"
	"time"
)

func computeTop3() []string {
	fmt.Println("считаю… (дорого)")
	time.Sleep(200 * time.Millisecond)
	return []string{"apple", "banana", "cherry"}
}

var top3 = sync.OnceValue(computeTop3)

func main() {
	for i := range 8 {
		go func(id int) {
			fmt.Printf("горутина %d: %v\n", id, top3())
		}(i)
	}

	time.Sleep(time.Second)
}
