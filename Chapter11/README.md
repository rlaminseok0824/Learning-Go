# What I Learn

## 1
  ### 벤치마킹
  ``` shell
  go test -bench=. *.go > old.txt
  ```
  ### benchstat
  benchmark 한 내용의 둘을 비교하기 위한 툴
  ``` shell
    go install golang.org/x/perf/cmd/benchstat
  ```

  실행방법
  ``` shell
    ~/go/bin/benchstat [파일1] [파일2]
  ```