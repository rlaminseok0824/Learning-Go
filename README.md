# Learning-Go
Learning Go with Go 마스터하기


## GC(Garbage Collecter)

### 삼색 마크 앤 스윕 알고리즘(tricolor mark-and-sweep algorithm)
 * 핵심 원리
    * 힙에 있는 오브젝트를 검은색(흰색 집합 포인터 x), 회색(흰색 집합 중 일부), 흰색(검은색 o)의 세 가지 색깔로 지정된 집합으로 나눔
  모든 오브젝트 흰색 -> 루트 오브젝트(스택에 있는 오브젝트 or 전역변수처럼 직접 접근할 수 있는 오브젝트) 모두 회색으로
  gc는 회색을 검은색으로 바꾸고 if 흰색 오브젝트 가지고 있으면 회색에 넣고 최종적으로 회색 집합이 모두 없어질 때까지

### 마크 앤 스윕 알고리즘
    우선 프로그램의 실행을 모두 멈추고 힙에서 접근할 수 있는 오브젝트를 모두 방문한 뒤에 적절히 표시
    핵심 원리 불변이 되지 않게 하기 위하여 새 오브젝트는 모두 회색 집합으로 보냄 => 동작 중 실행 가능

java는 가비지 컬렉터 다양함 G1 가비지 컬렉터는 저지연 애플리케이션에 적합해 자주 사용 / But Go는 아니다.