// Задача 1. «Первый — победил»

// **Условие:** Сервису приходит 5 «запросов» (5 горутин). Приз один — забирает его только первый. Объявите `var prize atomic.Bool` и пусть каждая горутина «претендует» на приз вызовом `prize.CompareAndSwap(false, true)`; строку `Горутина N выигрывает` печатает **только** та горутина, чей вызов вернул `true`. После паузы `time.Sleep(time.Second)` программа завершается.

// **Пример:**
// ```
// Горутина 3 выигрывает

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
