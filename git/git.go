package git

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/banananbo/gomemo/config"
)

// ExecuteGitCommand runs a git command in the specified directory and handles standard output and errors
func ExecuteGitCommand(rootDir string, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = rootDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// PushMemo handles the complete process of pushing a memo
func PushMemo(config *config.Config) {
	fmt.Println("Pushing the memo...")

	// Check if the root directory exists
	if _, err := os.Stat(config.RootDir); os.IsNotExist(err) {
		fmt.Printf("Root directory %s does not exist\n", config.RootDir)
		return
	}

	// Stage all changes
	if err := ExecuteGitCommand(config.RootDir, "add", "."); err != nil {
		fmt.Println("Error adding changes:", err)
		return
	}

	// Commit the changes (optional)
	if err := ExecuteGitCommand(config.RootDir, "commit", "-m", "Auto commit from memo app"); err != nil && err.Error() != "exit status 1" {
		fmt.Println("Error committing changes:", err)
		return
	}

	// Push the changes to the remote repository
	if err := ExecuteGitCommand(config.RootDir, "push", "origin", "main"); err != nil {
		fmt.Println("Error pushing memo:", err)
	}
}
