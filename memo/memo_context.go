package memo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/banananbo/gomemo/config"
)

// MemoMode インターフェース
type MemoMode interface {
	DetermineMemoLocation(config *config.Config, category *string) (string, string)
	CreateNewFile(filepath string) error
}

// MemoContext 構造体（戦略を持つコンテキスト）
type MemoContext struct {
	mode MemoMode
}

// SetMode メソッド（戦略を動的に設定）
func (c *MemoContext) SetMode(mode MemoMode) {
	c.mode = mode
}

// OpenMemo メソッド（戦略に基づいてメモを開く）
func (c *MemoContext) OpenMemo(config *config.Config, category *string) {
	memoDir, filename := c.mode.DetermineMemoLocation(config, category)
	if memoDir == "" || filename == "" {
		return
	}

	filepath := filepath.Join(memoDir, filename)

	// ディレクトリが存在しない場合は作成
	if err := os.MkdirAll(memoDir, os.ModePerm); err != nil {
		fmt.Printf("ディレクトリの作成に失敗しました: %v\n", err)
		return
	}

	// ファイルが存在しない場合は作成
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		if err := c.mode.CreateNewFile(filepath); err != nil {
			fmt.Println("ファイルの作成中にエラーが発生しました:", err)
			return
		}
	}

	fmt.Printf("メモが作成されました: %s\n", filepath)

	openInVim(filepath)
}
