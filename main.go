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
		createMemo(config, "")
		return
	}

	command := os.Args[1]

	switch command {
	case "life":
		// "life" モードでメモを作成
		createMemo(config, "life")

	case "push":
		// pushCommand 実行
		pushCommand(config)

	default:
		// "cat=xx" のようなカテゴリ指定の引数を解析
		if strings.HasPrefix(command, "cat=") {
			// "=" 以降をカテゴリ名として取得
			category := strings.TrimPrefix(command, "cat=")
			createMemo(config, category)
		} else {
			// 不明なコマンドは通常のメモ作成
			createMemo(config, "")
		}
	}
}

// Function to handle memo creation
func createMemo(config *config.Config, modeOrCategory string) {
	var memoDir string = config.DefaultMemoDir
	var filename string

	if modeOrCategory == "life" {
		fmt.Println("Creating a memo in 'life' mode...")
		memoDir = config.LifeMemoDir
		filename = modeOrCategory + ".md"
	} else if modeOrCategory != "" {
		fmt.Printf("Creating a memo in category '%s'...\n", modeOrCategory)
		memoDir = config.CategoriesDir
		filename = modeOrCategory + ".md"
	} else {
		fmt.Println("Creating a default memo...")
		// デフォルトのメモ作成
		currentTime := time.Now()
		filename = currentTime.Format("20060102") + ".txt"
	}

	fmt.Println("Value pf filename:", filename)
	fmt.Println("Value pf memodir:", memoDir)

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
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error opening file with vim:", err)
	}
}

// Function to handle memo pushing
func pushCommand(config *config.Config) {
	// Add your push logic here, such as git push or any other action needed
	fmt.Println("Pushing the memo...") // Placeholder for actual push logic

	// Check if the root directory exists
	if _, err := os.Stat(config.RootDir); os.IsNotExist(err) {
		fmt.Printf("Root directory %s does not exist\n", config.RootDir)
		return
	}

	addCmd := exec.Command("git", "add", ".") // Stage all changes
	addCmd.Dir = config.RootDir               // Set the working directory to RootDir
	addCmd.Stdout = os.Stdout
	addCmd.Stderr = os.Stderr
	err := addCmd.Run()
	if err != nil {
		fmt.Println("Error adding changes:", err)
		return
	}

	// Commit the changes (optional, if you want to commit as well)
	commitCmd := exec.Command("git", "commit", "-m", "Auto commit from memo app") // Commit changes with a message
	commitCmd.Dir = config.RootDir
	commitCmd.Stdout = os.Stdout
	commitCmd.Stderr = os.Stderr
	err = commitCmd.Run()
	if err != nil {
		// If there's nothing to commit, continue with git push
		if err.Error() != "exit status 1" {
			fmt.Println("Error committing changes:", err)
			return
		}
	}

	// Push the changes to the remote repository
	pushCmd := exec.Command("git", "push", "origin", "main") // Push to the main branch
	pushCmd.Dir = config.RootDir
	pushCmd.Stdout = os.Stdout
	pushCmd.Stderr = os.Stderr
	err = pushCmd.Run()
	if err != nil {
		fmt.Println("Error pushing memo:", err)
		return
	}
}
