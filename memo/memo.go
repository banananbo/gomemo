package memo

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/banananbo/gomemo/config"
)

// Function to handle memo creation
func CreateMemo(config *config.Config, mode string, category *string) {
	var memoDir string
	var filename string

	switch mode {
	case "life":
		fmt.Println("Creating a memo in 'life' mode...")
		memoDir = config.LifeMemoDir
		currentTime := time.Now()
		filename = currentTime.Format("200601") + ".md"

	case "cat":
		if category != nil {
			fmt.Printf("Creating a memo in category '%s'...\n", *category)
			memoDir = config.CategoriesDir
			filename = *category + ".md"
		} else {
			fmt.Println("Category not specified for 'cat' mode")
			return
		}

	case "code":
		if category != nil {
			fmt.Printf("Creating a memo in code category '%s'...\n", *category)
			memoDir = config.CodesDir
			filename = *category + ".md"
		} else {
			fmt.Println("Category not specified for 'code' mode")
			return
		}

	default:
		// Default memo creation with date as filename
		fmt.Println("Creating a default memo...")
		currentTime := time.Now()
		memoDir = config.DefaultMemoDir
		filename = currentTime.Format("20060102") + ".md"
	}

	filepath := filepath.Join(memoDir, filename)

	// Ensure the directory exists
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
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error opening file with vim:", err)
	}
}
