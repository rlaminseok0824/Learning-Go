package main

import "fmt"

//경쟁상태 확인 방법 : go run -race chRace.go

func printer(ch chan<- bool, times int) {
	//channel에 쓰고 바로 close를 하기 때문에 일이 순차적으로 진행이 된다.
	for i := 0; i < times; i++ {
		ch <- true
	}
	close(ch)
}

func main() {
	//buffer 없이 chan을 만든다
	var ch chan bool = make(chan bool)

	go printer(ch, 5)

	//앞서 go-routine 때 close()로 인하여 ch가 닫혔기 때문에 range까지 가면 스스로 닫힌다.
	//따라서 경쟁상태 발생 x
	for val := range ch {
		fmt.Print(val, " ")
	}
	fmt.Println()

	for i := 0; i < 15; i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println()
}
