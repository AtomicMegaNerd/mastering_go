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
		return
	}

	for ix := 1; ix != len(arguments); ix++ {
		file := arguments[ix]
		path := os.Getenv("PATH")
		// This splits the path into a list of strings so we can search each
		pathSplit := filepath.SplitList(path)
		// For each directory in the list of paths
		for _, directory := range pathSplit {
			// Get the full path
			fullPath := filepath.Join(directory, file)
			// See if the file exists at all
			fileInfo, err := os.Stat(fullPath)
			if err == nil {
				// The file exists but it is a directory or a symlink?
				mode := fileInfo.Mode()
				// Is it a regular file?
				if mode.IsRegular() {
					if mode&0111 != 0 {
						fmt.Println(fullPath)
					}
				}
			}
		}
	}
}
