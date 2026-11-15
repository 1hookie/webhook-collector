package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var score atomic.Int64

func main() {
	for range 20 {
		go func() {
			for range 100 {
				score.Add(1)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("score:", score.Load())
}
