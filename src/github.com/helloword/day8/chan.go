package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//写方法
func write(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 50)
	}

	close(ch)
	wg.Done()
}

//读方法
func read(ch chan int) {
	for v := range ch {
		fmt.Println("读取到数据", v)
	}
	wg.Done()
}

func main() {
	var ch = make(chan int, 100)

	wg.Add(1)
	write(ch)
	wg.Add(1)
	read(ch)
}
