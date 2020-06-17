package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch1 <- i
	}

	ch2 := make(chan string, 10)
	for i := 0; i < 10; i++ {
		ch2 <- "hello" + fmt.Sprintf("%v", i)
	}

	for {
		select {
		case v := <-ch1:
			fmt.Printf("从ch1中读取的数据%v\n", v)
			time.Sleep(time.Millisecond * 50)
		case v := <-ch2:
			fmt.Printf("从ch2中读取的数据%v\n", v)
			time.Sleep(time.Millisecond * 50)
		default:
			fmt.Println("数据获取完毕")
			return
		}
	}
}
