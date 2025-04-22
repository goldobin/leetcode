package print_foobar_alternately

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_FooBar(t *testing.T) {
	const n = 1000
	var (
		m        sync.Mutex
		wg       sync.WaitGroup
		calls    []bool
		printFoo = func() {
			m.Lock()
			defer m.Unlock()
			calls = append(calls, true)
		}
		printBar = func() {
			m.Lock()
			defer m.Unlock()
			calls = append(calls, false)
		}
		fb  = NewFooBar(n)
		foo = func() {
			defer wg.Done()
			fb.Foo(printFoo)
		}
		bar = func() {
			defer wg.Done()
			fb.Bar(printBar)
		}
	)

	wg.Add(2)

	go foo()
	go bar()
	wg.Wait()

	require.Equal(t, n*2, len(calls))
	for i := 0; i < n*2; i += 2 {
		assert.Equal(t, true, calls[i])
		assert.Equal(t, false, calls[i+1])
	}
}
