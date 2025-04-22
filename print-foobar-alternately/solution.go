package print_foobar_alternately

import "sync"

type FooBar struct {
	cond *sync.Cond
	flag bool
	n    int
}

func NewFooBar(n int) *FooBar {
	result := &FooBar{
		cond: sync.NewCond(&sync.Mutex{}),
		flag: true,
		n:    n,
	}

	return result
}

func (fb *FooBar) Foo(printFoo func()) {
	c := fb.cond
	for i := 0; i < fb.n; i++ {
		c.L.Lock()
		for !fb.flag {
			c.Wait()
		}
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.flag = false
		c.Signal()
		c.L.Unlock()
	}
}

func (fb *FooBar) Bar(printBar func()) {
	c := fb.cond
	for i := 0; i < fb.n; i++ {
		c.L.Lock()
		for fb.flag {
			c.Wait()
		}
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		fb.flag = true
		c.Signal()
		c.L.Unlock()
	}
}
