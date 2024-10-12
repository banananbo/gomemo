package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/banananbo/gomemo/config"
)

func main() {
	// Load the config file
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	var memoDir string = config.DefaultMemoDir
	var filename string

	// Perse command line arguments
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--life":
			// --life オプションの場合、LifeMemoDir を使用
			memoDir = config.LifeMemoDir
		case "-c":
			memoDir = config.CategoriesDir
			// -c オプションでカテゴリ名を指定
			if i+1 < len(os.Args) {
				category := os.Args[i+1]
				filename = category + ".txt" // use category name as filename
				i++                          // skip the next argument
			}
		}
	}

	// If filename is not specified, use the current date
	if filename == "" {
		currentTime := time.Now()
		filename = currentTime.Format("20060102") + ".txt"
	}

	filepath := filepath.Join(memoDir, filename)

	// If the directory does not exist, create it
	if err := os.MkdirAll(memoDir, os.ModePerm); err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		return
	}

	// If file does not exist, create it
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		file, err := os.Create(filepath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		file.Close()
	}

	// Open the file with vim
	cmd := exec.Command("vim", "-c", "normal Go", "-c", "startinsert", filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error opening file with vim:", err)
	}
}
