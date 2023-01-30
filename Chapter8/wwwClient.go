package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if (len(os.Args)) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	URL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error in parsing: ", err)
		return
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	request, err := http.NewRequest(http.MethodGet, URL.String(), nil) //URL의 Request 객체 반환
	if err != nil {
		fmt.Println("Get: ", err)
		return
	}

	httpData, err := c.Do(request) //http.Response 반환
	if err != nil {
		fmt.Println("Error in Do():", err)
		return
	}

	fmt.Println("Status Code: ", httpData.Status)
	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Print(string(header))

	contentType := httpData.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("Character Set : ", characterSet[1])
	}

	if httpData.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength:", httpData.ContentLength)
	}

	//http 응답 크기를 구하기 위한 방법
	length := 0
	var buffer [1024]byte
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length = length + n
	}
	fmt.Println("Calculated response data length:", length)
}
