// Задача 2. «Уникальная регистрация»
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	registered atomic.Int64
	words      sync.Map // слово -> true (зарегистрировано)
)

// Register: true, если слово зарегистрировано впервые.
func Register(w string) bool {
	_, loaded := words.LoadOrStore(w, true)
	return !loaded
}

func main() {
	for i := range 5 {
		go func(id int) {
			for j := range 20 {
				w := fmt.Sprintf("w%d", j)
				if Register(w) {
					registered.Add(1)
				}
			}
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("registered:", registered.Load())
}
