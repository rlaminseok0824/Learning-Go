package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(key string) (*Entry, error) {
	for i, v := range data {
		if v.Surname == key {
			return &data[i], nil
		}
	}
	return nil, nil
}

func list() error {
	for _, v := range data {
		fmt.Println(v)
	}
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	data = append(data, Entry{"Mihalis", "Tsoukalos", "2198416471"})
	data = append(data, Entry{"Mary", "Doe", "2198416471"})
	data = append(data, Entry{"John", "Black", "2198416471"})

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result, _ := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)
	case "list":
		list()
		return
	default:
		fmt.Println("Not a valid option")
	}
	return
}
