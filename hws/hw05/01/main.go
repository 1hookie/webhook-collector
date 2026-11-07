// Задача 1. «Первый — победил»
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var prize atomic.Bool
	for i := range 5 {
		go func(id int) {
			if prize.CompareAndSwap(false, true) {
				fmt.Printf("Горутина %d выигрывает", id)
			}
		}(i)
	}

	time.Sleep(time.Second)
}
