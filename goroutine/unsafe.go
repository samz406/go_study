package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var wg sync.WaitGroup

var mutex sync.Mutex

var rw sync.RWMutex

func add()  {
	count++
	fmt.Println("this count is ", count)
	time.Sleep(time.Minute*2)
	wg.Done()
}

func lockAdd()  {

	mutex.Lock()
	count++
	fmt.Println("this count is ", count)
	time.Sleep(time.Millisecond)
	wg.Done()
	mutex.Unlock()
}

// 读写锁
func read()  {
	rw.RLock()
	fmt.Println("read opt")
	time.Sleep(time.Second*2)
	wg.Done()
	rw.RUnlock()
}

func write()  {
	rw.Lock()
	fmt.Println("write opt")
	time.Sleep(time.Second*2)
	wg.Done()
	rw.Unlock()

}


func main() {
	//for i := 0; i < 1000*1000; i++ {
	//	wg.Add(1)
	//	//go lockAdd()
	//	go add()
	//}
	//wg.Wait()

	//read write lock
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
}


