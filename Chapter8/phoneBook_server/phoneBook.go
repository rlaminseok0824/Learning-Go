package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

// 전화번호에 들어갈 목록을 struct로 선언
type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

// CSVFILE이 저장될 위치를 전역변수로 선언함
var CSVFILE = "./data.csv"

// Entry의 slice를 PhoneBook으로 함 == typedef
type PhoneBook []Entry

// data를 phoneBook의 인터페이스로 설정
var data = PhoneBook{}

// search 및 delete할 때의 key를 찾기위한 index 만듬
var index map[string]int

// 처음 시작할 때 읽을 readCSVFile을 위한 함수
func readCSVFile(filepath string) error {
	_, err := os.Stat(filepath) //해당 file이 있는지 찾음
	if err != nil {
		return err
	}

	f, err := os.Open(filepath) // 있다면 file을 open한다
	if err != nil {
		return err
	}
	defer f.Close() //defer를 통해 함수가 끝나면 자동적으로 Close()하게 만든다.

	lines, err := csv.NewReader(f).ReadAll() // 열었던 file(f)의 전체를 읽는다.
	if err != nil {
		return err
	}

	for _, line := range lines { //이 lines 의 range만큼 _ = index, line = value
		temp := Entry{ //csv 파일엔 ,를 기점으로 나뉘어져있음 => index로 구분 가능
			Name:       line[0],
			Surname:    line[1],
			Tel:        line[2],
			LastAccess: line[3],
		}
		//이를 data interface에 추가함
		data = append(data, temp)
	}
	return nil
}

// File에 저장하기 위한 함수
func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath) //새로 filePath를 만든다. 있는지 없는지 구분 x
	if err != nil {
		return err
	}
	defer csvfile.Close() //defer를 통해 함수가 끝나면 자동적으로 닫게 함

	csvwriter := csv.NewWriter(csvfile) //NewReader와 다른 쓰기 위한 NewWriter
	for _, row := range data {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastAccess} //slice 선언 및 값을 넣음
		_ = csvwriter.Write(temp)                                        //temp를 파일에 씀
	}
	csvwriter.Flush() //Flush를 통해 파일에 최종적으로 저장
	return nil
}

func createIndex() error {
	index = make(map[string]int) // index는 string : int의 dict 구조
	// Tel을 key로 자료의 index를 구할 수 있음
	for i, k := range data {
		key := k.Tel
		index[key] = i
	}
	return nil
}

// 자동적으로 LastAccess 갱신해줌
func initS(N, S, T string) *Entry {
	if T == "" || S == "" {
		return nil
	}

	LastAccess := strconv.FormatInt(time.Now().Unix(), 10) //time을 10진수의 string 형태로 지정해주는 함수
	return &Entry{Name: N, Surname: S, Tel: T, LastAccess: LastAccess}
}

// insert Entry의 주소값을 받음? => Entry의 크기가 커질 수 있기 때문에
func insert(pS *Entry) error {
	// If it already exists, do not add it
	_, ok := index[(*pS).Tel] //pS가 포인터로 받았으므로 pS의 실제값인 (*pS)의 Tel값
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}

	//앞의 (*pS) == *&pS 와 같은 의미
	*&pS.LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	data = append(data, *pS)
	// Update the index
	_ = createIndex()

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

// key == Tel을 기준으로 delete함
func deleteEntry(key string) error {
	i, ok := index[key] //index에 key가 존재하는 지 파악
	if !ok {            //if key가 존재하지 않다면 error return
		return fmt.Errorf("%s cannot be found!", key)
	}
	//i번째 index를 삭제하는 방법, [:i]는 0~i-1번까지를 slice에 저장
	data = append(data[:i], data[i+1:]...)
	// Update the index - key does not exist any more
	// key를 가진 index 삭제
	delete(index, key)

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

// key = Tel 을 통해 search함. delete와 유사
func search(key string) *Entry {
	i, ok := index[key]
	if !ok {
		return nil
	}

	data[i].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	return &data[i]
}

func matchTel(s string) bool {
	t := []byte(s)                   //정규표현식을 match하기 위해 string -> byte로 변환
	re := regexp.MustCompile(`\d+$`) //정규표현 매치
	return re.Match(t)
}

func list() string {
	var all string
	for _, k := range data {
		all = all + k.Name + " " + k.Surname + " " + k.Tel + "\n"
	}
	return all
}

func main() {
	err := readCSVFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println("Cannot create index.")
		return
	}

	mux := http.NewServeMux() //서버 멀티플렉싱 해줌 => HandleFunc을 지원해줌
	s := &http.Server{        //서버를 열기 위해 만족할 struct
		Addr:         PORT, // IP주소
		Handler:      mux,  //각 구현마다 handlerFunc위해
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	//각 Handle 구현 가능
	mux.Handle("/list", http.HandlerFunc(listHandler))
	mux.Handle("/insert/", http.HandlerFunc(insertHandler))
	mux.Handle("/insert", http.HandlerFunc(insertHandler))
	mux.Handle("/search", http.HandlerFunc(searchHandler))
	mux.Handle("/search/", http.HandlerFunc(searchHandler))
	mux.Handle("/delete/", http.HandlerFunc(deleteHandler))
	mux.Handle("/status", http.HandlerFunc(statusHandler))
	mux.Handle("/", http.HandlerFunc(defaultHandler))

	fmt.Println("Ready to serve at", PORT)
	//http 서버 시작
	err = s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
