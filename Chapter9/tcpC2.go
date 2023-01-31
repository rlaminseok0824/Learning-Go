package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a server:port string!")
		return
	}

	connect := arguments[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", connect) //오직 tcp 통신에서만 사용가능 // IPv4를 사용하고 TCP address의 end point를 반환
	if err != nil {
		fmt.Println("ResolveTCPAdder: ", err)
		return
	}

	conn, err := net.DialTCP("tcp4", nil, tcpAddr) // tcp addr의 end point를 가지고 tcp 연결 시도, 두번째 파라미터가 nil == localhost
	if err != nil {
		fmt.Println("DialTCP: ", err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
