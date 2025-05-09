package cond_broadcast

import (
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
		b.Subscribe(func() {
			defer clicked.Done()
			counter.Add(1)
		})
	}

	b.Clicked.Broadcast()
	clicked.Wait()
	return int(counter.Load())
}
