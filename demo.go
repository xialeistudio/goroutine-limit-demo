package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan int, 100)
	for i := 0; i <= 1000; i++ {
		ch <- 1
		go worker(i, ch)
	}
	close(ch)
	log.Println("complete", time.Since(start).Seconds())
}

// 模拟耗时操作
func worker(i int, ch chan int) {
	log.Println("worker", i)
	time.Sleep(time.Second)
	<-ch
}
