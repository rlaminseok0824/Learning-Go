package main

import (
	"fmt"
	"time"
)

func printme(x int) {
	fmt.Println("*", x)
	return
}

func main() {
	//go + 익명함수를 통하여 go-rutine 실행 가능(쓰레드) 뒤에 매개변수 10을 통해 입력값 10 전달
	go func(x int) {
		fmt.Printf("%d ", x)
	}(10)
	//다음과 같이 앞서 정의한 함수를 바로 go-routine으로 실행 가능
	go printme(15)

	// but 순서대로 10 * 15 가 나와야할 것 같지만 go-routine 특성상 어느 것이 먼저 끝날지 모르므로 15 * 10 등 이상하게 나옴
	time.Sleep(time.Second)
	fmt.Println("Exiting....")
}
