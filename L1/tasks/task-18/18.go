package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	sync.Mutex
	count int64
}

func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	atomic.AddInt64(&c.count, 1)
}

func (c *Counter) GetCount() int64 {
	return atomic.LoadInt64(&c.count)
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Итоговый счетчик:", counter.GetCount())
}
