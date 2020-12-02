package main

import (
	"fmt"
	"time"
)

func main() {
	//管道
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20   //插入
	m1 := <-ch1 //读取
	m2 := <-ch1
	println("channel print out ", m1, m2)

	ch2 := make(chan<- int, 2)
	ch2 <- 10
	ch2 <- 20

	//m3 := <- ch2 //error Invalid operation: <- ch2 (receive from send-only type chan<- int)

	//select

	intChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		intChan <- i
		time.Sleep(time.Second * 2)
	}

	stringChan := make(chan string, 10)

	for i := 0; i < 10; i++ {

		stringChan <- "the number is " + fmt.Sprintf("%d", i)
	}

	//select 不需要close channel
	for {
		select {
		case i := <-intChan:
			fmt.Printf("the intChannel print out is %d\n", i)
		case i := <-stringChan:
			fmt.Printf("the stringChannel print out is %s\n", i)
		default:
			fmt.Printf("print complete\n")
			return
		}
	}

}
