package memo

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func createFromTemplate(memoPath string, templatePath string) error {
	// テンプレートを読み込む
	templateFile, err := os.Open(templatePath)
	if err != nil {
		return fmt.Errorf("テンプレートの読み込みに失敗しました: %v", err)
	}
	defer templateFile.Close()

	// 新しいメモファイルを作成
	newFile, err := os.Create(memoPath)
	if err != nil {
		return fmt.Errorf("メモファイルの作成に失敗しました: %v", err)
	}
	defer newFile.Close()

	// テンプレート内容を新しいファイルにコピー
	_, err = io.Copy(newFile, templateFile)
	if err != nil {
		return fmt.Errorf("テンプレートからファイルへのコピー中にエラーが発生しました: %v", err)
	}

	return nil
}

func createEmptyFile(memoPath string) error {
	file, err := os.Create(memoPath)
	if err != nil {
		return fmt.Errorf("ファイルの作成に失敗しました: %v", err)
	}
	defer file.Close()
	return nil
}

// vimでファイルを開くヘルパー関数
func openInVim(filepath string) {
	cmd := exec.Command("vim", "-c", "normal Go", "-c", "startinsert", filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("vimでファイルを開く際にエラーが発生しました:", err)
	}
}
