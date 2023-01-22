package main

import (
	"fmt"
	"sync"
	"time"
)

// Golang에서는 mutex가 하나의 lock만을 사용가능한 한계를 극복하기 위해 RWMutex를 지원
// 따라서 다중 읽기가 가능해지지만 하나라도 Lock이 된 상태에서 Write가 불가능해진다.

var Password *secret
var wg sync.WaitGroup

type secret struct {
	RWM      sync.RWMutex
	password string
}

func Change(pass string) {
	fmt.Println("Change() function")
	Password.RWM.Lock()
	fmt.Println("Change() Locked")
	time.Sleep(4 * time.Second)
	Password.password = pass
	Password.RWM.Unlock()
	fmt.Println("Change() UnLocked")
}

func show() {
	defer wg.Done()
	Password.RWM.RLock()
	fmt.Println("Show function locked!")
	time.Sleep(2 * time.Second)
	fmt.Println("Pass value:", Password.password)
	defer Password.RWM.RUnlock()
}

func main() {
	Password = &secret{password: "myPass"}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go show()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change("123456")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change("54321")
	}()

	wg.Wait()

	// Direct access to Password.password
	fmt.Println("Current password value:", Password.password)
}
