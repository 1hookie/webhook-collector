package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Player struct {
	Name  string
	Score int
}

var leader atomic.Pointer[Player]

func tryWin(p *Player) {
	for cur := leader.Load(); cur == nil || p.Score > cur.Score; {
		if leader.CompareAndSwap(cur, p) {
			return
		}
		leader.Load()
	}
}

func main() {
	for i := 1; i <= 10; i++ {
		p := &Player{Name: fmt.Sprintf("p%d", i), Score: i * 10}
		go func() {
			tryWin(p)
		}()
	}

	time.Sleep(time.Second)
	l := leader.Load()
	if l == nil {
		fmt.Println("leader: (пусто) — заполните tryWin")
		return
	}
	fmt.Println("leader:", l.Name, l.Score)
}
