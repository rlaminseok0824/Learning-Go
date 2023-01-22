package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int
}

type Result struct {
	job    Client
	square int
}

var size = runtime.GOMAXPROCS(0) //go-routine의 수 즉, 쓰레드의 수 지정
var clients = make(chan Client, size)
var data = make(chan Result, size)

func worker(wg *sync.WaitGroup) {
	for c := range clients {
		square := c.integer * c.integer
		output := Result{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	wg.Done()
}

func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need #jobs and #workers!")
		return
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	nWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	go create(nJobs)

	finished := make(chan interface{}) //프로그램 실행을 블록하기 위한 용도로 사용 => 타입 지정 필요 x
	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\tint: ", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
		}
		finished <- true
	}()

	var wg sync.WaitGroup
	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(data)

	fmt.Printf("Finished: %v\n", <-finished)
}
