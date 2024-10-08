package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"memo/config"
)

func main() {
	// 設定ファイルを読み込む
	config, err := config.LoadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	var memoDir string
	var filename string

	// コマンドライン引数の解析
	if len(os.Args) > 1 && os.Args[1] == "--life" {
		memoDir = config.LifeMemoDir
		if len(os.Args) > 2 {
			filename = os.Args[2]
		}
	} else {
		memoDir = config.DefaultMemoDir
		if len(os.Args) > 1 {
			filename = os.Args[1]
		}
	}

	// ファイル名が指定されていない場合は、日付をファイル名にする
	if filename == "" {
		currentTime := time.Now()
		filename = currentTime.Format("20060102") + ".txt"
	}

	filepath := filepath.Join(memoDir, filename)

	// ファイルが存在しない場合は作成
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		file, err := os.Create(filepath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		file.Close()
	}

	// Vimでファイルを末尾から入力モードで開く
	cmd := exec.Command("vim", "-c", "normal Go", "-c", "startinsert", filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error opening file with vim:", err)
	}
}
