package git

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/banananbo/gomemo/config"
)

// Function to add changes to the Git staging area
func AddChanges(rootDir string) error {
	addCmd := exec.Command("git", "add", ".")
	addCmd.Dir = rootDir
	addCmd.Stdout = os.Stdout
	addCmd.Stderr = os.Stderr
	return addCmd.Run()
}

// Function to commit changes to the Git repository
func CommitChanges(rootDir string, message string) error {
	commitCmd := exec.Command("git", "commit", "-m", message)
	commitCmd.Dir = rootDir
	commitCmd.Stdout = os.Stdout
	commitCmd.Stderr = os.Stderr
	return commitCmd.Run()
}

// Function to push changes to the remote Git repository
func PushChanges(rootDir string, branch string) error {
	pushCmd := exec.Command("git", "push", "origin", branch)
	pushCmd.Dir = rootDir
	pushCmd.Stdout = os.Stdout
	pushCmd.Stderr = os.Stderr
	return pushCmd.Run()
}

// Function to handle memo pushing
func PushMemo(config *config.Config) {
	fmt.Println("Pushing the memo...")

	// Check if the root directory exists
	if _, err := os.Stat(config.RootDir); os.IsNotExist(err) {
		fmt.Printf("Root directory %s does not exist\n", config.RootDir)
		return
	}

	// Stage all changes
	if err := AddChanges(config.RootDir); err != nil {
		fmt.Println("Error adding changes:", err)
		return
	}

	// Commit the changes (optional)
	if err := CommitChanges(config.RootDir, "Auto commit from memo app"); err != nil && err.Error() != "exit status 1" {
		fmt.Println("Error committing changes:", err)
		return
	}

	// Push the changes to the remote repository
	if err := PushChanges(config.RootDir, "main"); err != nil {
		fmt.Println("Error pushing memo:", err)
	}
}
