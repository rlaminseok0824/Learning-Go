package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("You are using ", runtime.Compiler, " ")    //무슨 컴파일러를 사용하고 있는지 알려준다.
	fmt.Println("on a", runtime.GOARCH, "machine")        // 어느 운영체제에서 운영되고 있는지 알려준다.
	fmt.Println("Using Go version", runtime.Version())    // 현재 사용중인 Golang 의 version을 알려준다.
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0)) //현재 GOMAXPROCS의 개수를 알려준다.

	//GOMAXPROCS이란 Golang에서 최대 스케줄링을 조절해줄 수 있는 환경 변수로 선언할 떄 0을 선언하면 현재의 값으로 설정하겠다는 뜻
	// 1보다 큰 숫자가 입력이 된다면 현재의 GOMAXPROCS의 갯수를 옮기겠다는 뜻이다.
}
