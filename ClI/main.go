package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	dirPath := flag.String("dir", ".", "Directory")
	flag.Parse()

	entries, err := os.ReadDir(*dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println("Dir:", entry.Name())
		}
		info, err := entry.Info()
		if err != nil {
			fmt.Println("Error getting file info:", err)
			continue
		}
		fmt.Printf("File: %s, Size: %d bytes\n", entry.Name(), info.Size())
	}
}
