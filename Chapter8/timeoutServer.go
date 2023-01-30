package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

//서버에서의 타임아웃 설정이 더 중요함!!!
// 서버에 너무 많은 수의 연결이 맺어져 있을 경우 그 연결들을 처리해야하기 때문에
// 소프트웨어 버그 및 Dos 공격이 될 수 있기 때문에

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}
	fmt.Println("Using port number: ", PORT)

	m := http.NewServeMux()
	srv := &http.Server{
		Addr:         PORT,
		Handler:      m,
		ReadTimeout:  3 * time.Second, //본문을 포함한 요청 전체를 읽는데 걸리는 최대 시간 설정
		WriteTimeout: 3 * time.Second, //응답을 보내는데 허용한 최대 시간 설정
	}

	m.HandleFunc("/time", timeHandler)
	m.HandleFunc("/", myHandler)

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
