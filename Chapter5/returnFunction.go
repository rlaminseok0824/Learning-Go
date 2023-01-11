package main

import "fmt"

func funRet(i int) func(int) int {
	if i < 0 {
		return func(k int) int {
			k = -k
			return k + k
		}
	}

	return func(k int) int {
		return k * k
	}
}

func main() {
	n := 10
	//처음 선언하였을 때의 값에 의해서 funRet의 조건을 결정하게된다. i = k*k / j = k+k로
	//이후 다시 값이 지정이 되었을 때 정해진 함수별로 return이 된다.

	i := funRet(n)
	// The -4 parameter is used for determining
	// the anonymous function that will be returned
	j := funRet(-4)

	fmt.Printf("%T\n", i)
	fmt.Printf("%T %v\n", j, j)
	fmt.Println("j", j, j(-5))

	// Same input parameter but DIFFERENT
	// anonymous functions assigned to i and j
	fmt.Println(i(10))
	fmt.Println(j(10))
}
