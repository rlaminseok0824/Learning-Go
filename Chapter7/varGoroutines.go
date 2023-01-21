package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

// 이전까지 time.Sleep을 통하여 임의의 시간을 기다려 goroutine이 기다릴 까지 기다렸지만 이번엔 다르다.
func main() {
	count := 10
	arguments := os.Args
	if len(arguments) == 2 {
		t, err := strconv.Atoi(arguments[1])
		if err == nil {
			count = t
		}
	}

	fmt.Printf("Going to create %d goroutines.\n", count)

	var waitGroup sync.WaitGroup
	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < count; i++ {
		waitGroup.Add(1) //경쟁상태를 피하기 위해서는 Add()를 호출해야한다.
		go func(x int) {
			defer waitGroup.Done() //defer는 함수의 끝에서 Done = Add(-1)
			fmt.Printf("%d ", x)
		}(i)
	}

	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait() //Wait이 반환된다면 다음으로 전환
	fmt.Println("\nExiting...")
}
