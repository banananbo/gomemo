package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/banananbo/gomemo/config"
	"github.com/banananbo/gomemo/git"
	"github.com/banananbo/gomemo/memo"
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
		memo.CreateMemo(config, "", nil)
		return
	}

	command := os.Args[1]

	switch command {
	case "life":
		// "life" mode
		memo.CreateMemo(config, "life", nil)

	case "push":
		// Push the memo
		git.PushMemo(config)

	default:
		// Handle "cat=xx" or "code=xx" commands
		if strings.HasPrefix(command, "cat=") {
			category := strings.TrimPrefix(command, "cat=")
			memo.CreateMemo(config, "cat", &category)
		} else if strings.HasPrefix(command, "code=") {
			category := strings.TrimPrefix(command, "code=")
			memo.CreateMemo(config, "code", &category)
		} else {
			// Unknown command, create default memo
			memo.CreateMemo(config, "", nil)
		}
	}
}
