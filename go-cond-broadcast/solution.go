package go_cond_broadcast

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Button struct {
	Clicked *sync.Cond
}

func (b *Button) Subscribe(fn func()) {
	running := sync.WaitGroup{}
	running.Add(1)

	go func() {
		running.Done()

		b.Clicked.L.Lock()
		defer b.Clicked.L.Unlock()
		b.Clicked.Wait()

		fn()
	}()

	running.Wait()
}

func solution(n int) int {
	b := Button{
		Clicked: sync.NewCond(&sync.Mutex{}),
	}

	clicked := sync.WaitGroup{}
	counter := atomic.Int32{}

	clicked.Add(n)

	for i := 0; i < n; i++ {
		actionNo := i + 1
		b.Subscribe(func() {
			defer clicked.Done()
			print(fmt.Sprintf("Action %d\n", actionNo))
			counter.Add(1)
		})
	}

	b.Clicked.Broadcast()
	clicked.Wait()
	return int(counter.Load())
}
