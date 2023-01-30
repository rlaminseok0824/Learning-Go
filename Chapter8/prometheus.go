package main

import (
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"runtime/metrics"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PORT is the TCP port number the server will listen to
var PORT = ":1234"

// goroutines 와 memory를 count할 gauge 만들기
var n_goroutines = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "packt",
		Name:      "n_goroutines",
		Help:      "Number of goroutines"})

var n_memory = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "packt",
		Name:      "n_memory",
		Help:      "Memory usage"})

func main() {
	rand.Seed(time.Now().Unix())
	//prometheus에 등록을 한다.
	prometheus.MustRegister(n_goroutines)
	prometheus.MustRegister(n_memory)
	//metrics에 등록하기 위한 두가지
	const nGo = "/sched/goroutines:goroutines"
	const nMem = "/memory/classes/heap/free:bytes"
	//metrics 생성
	getMetric := make([]metrics.Sample, 2)
	getMetric[0].Name = nGo
	getMetric[1].Name = nMem

	//prometheus를 handle하기 위해 /metrics에 handle해야한다.
	http.Handle("/metrics", promhttp.Handler())
	//main 함수에서는 서버를 / go-routine에서는 메트릭을 수집하는 용도로 사용해야한다.
	go func() {
		for { //영원히 실행하게 하는 역할
			for i := 1; i < 4; i++ { //메트릭의 값이 항상 변하게 만드는 함수
				go func() {
					_ = make([]int, 1000000)
					time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				}()
			}

			runtime.GC() //GOlang의 가비지 컬렉터를 실행하게 한다.
			metrics.Read(getMetric)
			goVal := getMetric[0].Value.Uint64()
			memVal := getMetric[1].Value.Uint64()
			time.Sleep(time.Duration(rand.Intn(15)) * time.Second)

			n_goroutines.Set(float64(goVal)) //metrics 수집
			n_memory.Set(float64(memVal))
		}
	}()

	log.Println("Listening to port", PORT)
	log.Println(http.ListenAndServe(PORT, nil))
}
