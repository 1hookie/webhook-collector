package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("я в горутине") // запустилась и пошла дальше, не ждём
	fmt.Println("я в main")
	time.Sleep(100 * time.Millisecond) // без этого main может завершиться раньше
}
