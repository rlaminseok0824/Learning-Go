package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

//client가 server로부터 정해진 시간안에 response를 받지 못하면 timeout이 되는 구현

var myUrl string
var delay int = 5
var wg sync.WaitGroup

type myData struct {
	r   *http.Response
	err error
}

// In packages that use contexts, convention is to pass them as
// the first argument to a function.
func connect(c context.Context) error {
	defer wg.Done()                              //시작하자마자 defer를 하여 Done을 할 수 있게 만든다
	data := make(chan myData, 1)                 //myData의 구조체를 가지는 data라는 채널 만듬
	tr := &http.Transport{}                      // 빈 인터페이스 생성
	httpClient := &http.Client{Transport: tr}    //변수 초기화
	req, _ := http.NewRequest("GET", myUrl, nil) //변수 초기화

	//서버와 통신하는 코드
	go func() {
		response, err := httpClient.Do(req)
		if err != nil {
			fmt.Println(err)
			data <- myData{nil, err}
			return
		} else {
			pack := myData{response, err}
			data <- pack
		}
	}()

	// select를 통하여 타임아웃 되었는지 실행할지 결정
	select {
	case <-c.Done(): //timeout 시
		tr.CancelRequest(req)
		<-data
		fmt.Println("The request was canceled!")
		return c.Err()
	case ok := <-data: // 정상인 경우
		err := ok.err
		resp := ok.r
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		defer resp.Body.Close()

		realHTTPData, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		// Although fmt.Printf() is used here, server processes
		// use the log.Printf() function instead.
		fmt.Printf("Server Response: %s\n", realHTTPData)
	}
	return nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need a URL and a delay!")
		return
	}

	myUrl = os.Args[1]
	if len(os.Args) == 3 {
		t, err := strconv.Atoi(os.Args[2]) // timeout 인자를 받으면 ascii to int로
		if err != nil {
			fmt.Println(err)
			return
		}
		delay = t //delay는 default로 5로 설정되어있음
	}

	fmt.Println("Delay:", delay)
	//context 사용 => withTimeout 사용하기 위해서
	c := context.Background()

	// context.withTimeout 앞서 설정한 delay로
	c, cancel := context.WithTimeout(c, time.Duration(delay)*time.Second)
	defer cancel()

	fmt.Printf("Connecting to %s \n", myUrl)
	//wait Group 사용 => 효율적인 go-routine 사용하기 위함
	wg.Add(1)
	go connect(c)
	wg.Wait()
	fmt.Println("Exiting...")
}
