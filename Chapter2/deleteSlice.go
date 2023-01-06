package main

import (
	"fmt"
	"os"
	"strconv"
)

// Go의 가장 핵심적인 기능 중 하나인 slice(유사 vector)에는 따로 delete기능이 구현이 되어있지않아 직접 구현하여야 한다.
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need an integer value.")
		return
	}

	index := arguments[1]
	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Using index", i)

	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", aSlice)

	// Delete element at index i
	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}

	//1. 삭제될 index 전후로 slice해서 다시 합치기
	// The ... operator auto expands aSlice[i+1:] so that
	// its elements can be appended to aSlice[:i] one by one
	aSlice = append(aSlice[:i], aSlice[i+1:]...)
	fmt.Println("After 1st deletion:", aSlice)

	// Delete element at index i
	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}

	//2. 삭제될 index에 맨 끝값 넣고 맨 끝값 슬라이스 하기(이게 더 좋아보임)
	// Replace element at index i with last element
	aSlice[i] = aSlice[len(aSlice)-1]
	// Remove last element
	aSlice = aSlice[:len(aSlice)-1]
	fmt.Println("After 2nd deletion:", aSlice)
}
