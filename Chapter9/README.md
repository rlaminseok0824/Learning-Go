# What I Learn

## netcat
### [관련 블로그](https://andrewpage.tistory.com/109)
TCP/UDP를 통해 연결하고, 데이터를 읽고 쓰는 유틸리티 프로그램
* 두 장비 간에 IP-Network 정상인지 확인

``` shell
    $ nc -l -p 9090 
```
-l : 서버로 구동
-p : port 바로 지정
-u : UDP로 설정(없으면 default : TCP)
-v/-vv : 출력 상세히

## 유닉스 도메인 socket

## 웹소켓
### [공식 github](https://github.com/gorilla/websocket)
