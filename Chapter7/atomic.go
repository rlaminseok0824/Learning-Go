package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//atomic package를 이용하면 간단한 더하기 연산같은 것을 mutex와 같은 복잡한 생각없이
//경쟁상태없이 구현 가능하지만 유연하게 사용이 불가하기 때문에 여전히 mutex는 필요로하다.

type atomCounter struct {
	val int64
}

func (c *atomCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	X := 100
	Y := 4
	var waitGroup sync.WaitGroup
	counter := atomCounter{}
	for i := 0; i < X; i++ {
		waitGroup.Add(1)
		go func(no int) {
			defer waitGroup.Done()
			for i := 0; i < Y; i++ {
				atomic.AddInt64(&counter.val, 1)
			}
		}(i)
	}

	waitGroup.Wait()
	fmt.Println(counter.Value())
}
