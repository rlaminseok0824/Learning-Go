package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second) //타이머 Second로 생성

	for {
		select {
		case input := <-c: //send로 만들어진 난수를 저장
			sum = sum + input
		case <-t.C: //if 시간초과면 c를 nil로 만들고 Done
			c = nil //chan이 nil이면 항상 블록됨 따라서 close() 다 필요없이 nil하면 된다.
			fmt.Println(sum)
			wg.Done()
		}
	}
}

func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func main() {
	c := make(chan int)
	rand.Seed(time.Now().Unix())

	wg.Add(1)
	go add(c)
	go send(c) //난수 생성
	wg.Wait()
}
