package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument")
	}
	files := arguments[1:]
	exeFile := make([]string, 0)

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, directory := range pathSplit {
		for _, file := range files {
			fullPath := filepath.Join(directory, file)
			// Does it exist?
			fileInfo, err := os.Stat(fullPath)
			if err == nil {
				mode := fileInfo.Mode()
				// Is it a regular file?
				if mode.IsRegular() {
					// Is it executable?
					if mode&0111 != 0 {
						fmt.Println(fullPath)
						exeFile = append(exeFile, fullPath)
						break
					}
				}
			}
		}

	}
}
