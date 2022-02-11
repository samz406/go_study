package main

//10个goroutine 同时计算
import (
	"fmt"
	"sync"
)

func main() {

	var count Counter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				count.Incr()
			}
		}()
	}
	wg.Wait()

	fmt.Println(count.get())
}

// 将添加逻辑封装到struc
type Counter struct {
	mu    sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) get() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
