package demo

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Start() {

	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()

	loadInt64 := atomic.LoadInt64(&counter)
	fmt.Println(loadInt64)
}
