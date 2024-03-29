package main

//10个goroutine 同时计算
import (
	"fmt"
	"sync"
)

func main() {

	var count = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	fmt.Println(count)
}
