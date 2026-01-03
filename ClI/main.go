package main

import (
	"flag" // For parsing command-line flags
	"fmt"
	"os" // For interacting with the filesystem
)

func main() {

	// Define a command-line flag "-dir" with default value "." (current directory)
	// flag.String returns a pointer to a string
	dirPath := flag.String("dir", ".", "Directory")

	// Parse the command-line flags
	flag.Parse()

	// Read all entries files and directories in the specified directory
	entries, err := os.ReadDir(*dirPath)
	if err != nil {
		// If there is an error reading the directory, print it and exit
		fmt.Println("Error reading directory:", err)
		return
	}

	// Loop through all directory entries
	for _, entry := range entries {

		// Check if the entry is a directory
		if entry.IsDir() {
			fmt.Println("Dir:", entry.Name()) // Print directory name
		}

		// Get more detailed info about the entry e.g size, permissions
		info, err := entry.Info()
		if err != nil {
			// If there is an error getting file info, print it and skip this entry
			fmt.Println("Error getting file info:", err)
			continue
		}

		// Print file name and size in bytes
		fmt.Printf("File: %s, Size: %d bytes\n", entry.Name(), info.Size())
	}
}
