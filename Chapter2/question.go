package main

import "fmt"

//Go 에서는 array는 쓰지 않는게 좋다.. 왜냐면 고정된 크기를 제공하기 때문 -> slice를 사용하면 훨씬 쉽게 구현 가능하다.

func MergeArrToSlice[T any](a, b [5]T) []T {
	newSlice := make([]T, 0)
	for _, v := range a {
		newSlice = append(newSlice, v)
	}

	for _, v := range b {
		newSlice = append(newSlice, v)
	}

	return newSlice
}

func MergeArrToArr[T any](a, b [5]T) [10]T {
	newArr := [10]T{}
	for i, v := range a {
		newArr[i] = v
	}
	for i, v := range b {
		newArr[i+5] = v
	}
	return newArr
}

func MergeSliceToArr[T any](a, b []T) [10]T {
	newArr := [10]T{}
	newSlice := make([]T, 10)
	copy(newSlice, a)
	copy(newSlice[5:], b)
	for i, v := range newSlice {
		newArr[i] = v
	}
	return newArr
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{6, 7, 8, 9, 10}

	c := MergeArrToSlice(a, b)
	fmt.Println("MergeArrToSlice : ", c)

	d := MergeArrToArr(a, b)
	fmt.Println("MergeArrToArr : ", d)

	e := []int{10, 9, 8, 7, 6}
	f := []int{6, 7, 8, 9, 10}
	fmt.Println("MergeSliceToArr : ", MergeSliceToArr(e, f))
}
