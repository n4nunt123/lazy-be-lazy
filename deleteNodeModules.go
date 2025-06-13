package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	targetDir := "your/target/directory"
	err := removeNodeModules(targetDir)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("All node_modules directories have been removed.")
	}
}

func removeNodeModules(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Walk error:", err)
			return nil
		}

		if info.IsDir() && info.Name() == "node_modules" {
			fmt.Println("Removing:", path)
			err := os.RemoveAll(path)
			if err != nil {
				fmt.Printf("Failed to remove %s: %v\n", path, err)
			} else {
				fmt.Printf("Removed: %s\n", path)
			}
			// Skip walking this directory’s children
			return filepath.SkipDir
		}
		return nil
	})
}
