package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

	if len(os.Args) < 2 {
		// No arguments passed, create default memo
		createMemo(config, "", nil)
		return
	}

	command := os.Args[1]

	switch command {
	case "life":
		// "life" mode
		createMemo(config, "life", nil)

	case "push":
		// Push the memo
		pushCommand(config)

	default:
		// Handle "cat=xx" or "code=xx" commands
		if strings.HasPrefix(command, "cat=") {
			category := strings.TrimPrefix(command, "cat=")
			createMemo(config, "cat", &category)
		} else if strings.HasPrefix(command, "code=") {
			category := strings.TrimPrefix(command, "code=")
			createMemo(config, "code", &category)
		} else {
			// Unknown command, create default memo
			createMemo(config, "", nil)
		}
	}
}

// Function to handle memo creation
func createMemo(config *config.Config, mode string, category *string) {
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

// Function to handle memo pushing
func pushCommand(config *config.Config) {
	fmt.Println("Pushing the memo...")

	// Check if the root directory exists
	if _, err := os.Stat(config.RootDir); os.IsNotExist(err) {
		fmt.Printf("Root directory %s does not exist\n", config.RootDir)
		return
	}

	// Stage all changes
	addCmd := exec.Command("git", "add", ".")
	addCmd.Dir = config.RootDir
	addCmd.Stdout = os.Stdout
	addCmd.Stderr = os.Stderr
	err := addCmd.Run()
	if err != nil {
		fmt.Println("Error adding changes:", err)
		return
	}

	// Commit the changes (optional)
	commitCmd := exec.Command("git", "commit", "-m", "Auto commit from memo app")
	commitCmd.Dir = config.RootDir
	commitCmd.Stdout = os.Stdout
	commitCmd.Stderr = os.Stderr
	err = commitCmd.Run()
	if err != nil && err.Error() != "exit status 1" { // Ignore "nothing to commit" errors
		fmt.Println("Error committing changes:", err)
		return
	}

	// Push the changes to the remote repository
	pushCmd := exec.Command("git", "push", "origin", "main")
	pushCmd.Dir = config.RootDir
	pushCmd.Stdout = os.Stdout
	pushCmd.Stderr = os.Stderr
	err = pushCmd.Run()
	if err != nil {
		fmt.Println("Error pushing memo:", err)
	}
}
