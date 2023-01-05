package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) <= 1 {
		fmt.Println("You need at least two arguments")
		return
	}
	invalid := make([]string, 0)
	var total, nInt, nFloat int
	for _, k := range arguments[1:] {
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInt++
			continue
		}

		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloat++
			continue
		}

		invalid = append(invalid, k)
	}

	fmt.Println("#read:", total, "#ints:", nInt, "#floats:", nFloat)
	if len(invalid) > total {
		fmt.Println("Too much invalid input:", len(invalid))
		for _, s := range invalid {
			fmt.Println(s)
		}
	}
}
