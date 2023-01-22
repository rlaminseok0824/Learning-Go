package main

import (
	"fmt"
	"time"
)

func printer(ch chan<- bool) {
	ch <- true
}

func writeToChannel(c chan<- int, x int) {
	fmt.Println("1", x)
	c <- x
	fmt.Println("2", x)
}

// <-chan은 읽기 전용, chan<- 쓰기 전용 미리 알려줌
func f2(out <-chan int, in chan<- int) {
	x := <-out
	fmt.Println("Read (f2):", x)
	in <- x
	return
}

func main() {
	c := make(chan int)
	go writeToChannel(c, 10) //channel의 사용 용도는 go-routine 동안에 계산값 저장위해
	time.Sleep(1 * time.Second)
	fmt.Println("Read:", <-c)
	time.Sleep(1 * time.Second)
	close(c) //사용 후 close 해주어야함

	c1 := make(chan int, 1)
	c2 := make(chan int, 1)

	c1 <- 5
	f2(c1, c2)
	fmt.Println("Read (main):", <-c2)
}
