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
				count.Lock()
				//临界区资源锁定
				count.Count++
				count.Unlock()
			}
		}()
	}
	wg.Wait()

	fmt.Println(count.Count)
}

type Counter struct {
	sync.Mutex
	Count uint64
}
