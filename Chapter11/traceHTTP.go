package main

import (
	"fmt"
	"net/http"
	"net/http/httptrace"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: URL\n")
		return
	}

	URL := os.Args[1]
	client := http.Client{}

	req, _ := http.NewRequest("GET", URL, nil)

	//http trace를 할 것들을 즐겨찾기 느낌 연결
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			fmt.Println("First response byte!")
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("Dial start")
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("Dial done")
		},
		WroteHeaders: func() {
			fmt.Println("Wrote headers")
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace)) //같이 감쌈 trace랑
	fmt.Println("Requesting data from server!")
	_, err := http.DefaultTransport.RoundTrip(req) //요청 추적하고자 하는거랑 같이 감쌈
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	// io.Copy(os.Stdout, response.Body)
}
