package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"golang.org/x/sync/semaphore"
)

// Maximum number of goroutines
var Workers = 4
var sem = semaphore.NewWeighted(int64(Workers)) // 새로운 thread(semaphore) 갯수 미리 지정

func worker(n int) int {
	square := n * n
	time.Sleep(time.Second)
	return square
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need #jobs!")
		return
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Where to store the results
	var results = make([]int, nJobs)

	// Needed by Acquire()
	ctx := context.TODO()

	for i := range results {
		err = sem.Acquire(ctx, 1) //Acquire를 통해 가중치를 정의가능함 현재는 1로 다 고정
		if err != nil {
			fmt.Println("Cannot acquire semaphore:", err)
			break
		}

		go func(i int) {
			defer sem.Release(1) // 모든 일을 마치고 Realease를 함
			temp := worker(i)
			// No race conditions here - each goroutine writes
			// to a different slice element
			results[i] = temp
		}(i)
	}

	// Acquire all of the tokens
	// This is similar to Wait()
	// It blocks until all workers have finished
	err = sem.Acquire(ctx, int64(Workers)) // Workers = 4 이기 때문에 앞서 가중치 1의 작업들이 다 끝나기 전까지 대기상태
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range results {
		fmt.Println(k, "->", v)
	}
}
