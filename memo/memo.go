package memo

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/banananbo/gomemo/config"
)

// MemoMode インターフェース
type MemoMode interface {
	DetermineMemoLocation(config *config.Config, category *string) (string, string)
	CreateNewFile(filepath string) error // 新しいメモファイルを作成するメソッド
}

// LifeMode 構造体（lifeモード用）
type LifeMode struct{}

// DetermineMemoLocation メソッド（lifeモード用の実装）
func (m LifeMode) DetermineMemoLocation(config *config.Config, category *string) (string, string) {
	currentTime := time.Now()
	filename := currentTime.Format("200601") + ".md"
	return config.LifeMemoDir, filename
}

// CreateNewFile メソッド（lifeモードではテンプレートを使用）
func (m LifeMode) CreateNewFile(memoPath string) error {
	templatePath := filepath.Join("templates", "life_template.md")
	return createFromTemplate(memoPath, templatePath)
}

// CatMode 構造体（catモード用）
type CatMode struct{}

// DetermineMemoLocation メソッド（catモード用の実装）
func (m CatMode) DetermineMemoLocation(config *config.Config, category *string) (string, string) {
	if category == nil {
		fmt.Println("「cat」モードに対するカテゴリーが指定されていません")
		return "", ""
	}
	return config.CategoriesDir, *category + ".md"
}

// CreateNewFile メソッド（catモードでもテンプレートを使用）
func (m CatMode) CreateNewFile(memoPath string) error {
	templatePath := filepath.Join("templates", "cat_template.md")
	return createFromTemplate(memoPath, templatePath)
}

// CodeMode 構造体（codeモード用）
type CodeMode struct{}

// DetermineMemoLocation メソッド（codeモード用の実装）
func (m CodeMode) DetermineMemoLocation(config *config.Config, category *string) (string, string) {
	if category == nil {
		fmt.Println("「code」モードに対するカテゴリーが指定されていません")
		return "", ""
	}
	return config.CodesDir, *category + ".md"
}

// CreateNewFile メソッド（codeモードではテンプレートを使用）
func (m CodeMode) CreateNewFile(memoPath string) error {
	templatePath := filepath.Join("templates", "code_template.md")
	return createFromTemplate(memoPath, templatePath)
}

// DefaultMode 構造体（デフォルトモード用）
type DefaultMode struct{}

// DetermineMemoLocation メソッド（デフォルトモード用の実装）
func (m DefaultMode) DetermineMemoLocation(config *config.Config, category *string) (string, string) {
	currentTime := time.Now()
	filename := currentTime.Format("20060102") + ".md"
	return config.DefaultMemoDir, filename
}

// CreateNewFile メソッド（デフォルトモードではテンプレートを使用しない）
func (m DefaultMode) CreateNewFile(memoPath string) error {
	return createEmptyFile(memoPath)
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

	memoPath := filepath.Join(memoDir, filename)

	// ディレクトリが存在しない場合は作成
	if err := os.MkdirAll(memoDir, os.ModePerm); err != nil {
		fmt.Printf("ディレクトリの作成に失敗しました: %v\n", err)
		return
	}

	// ファイルが存在しない場合は作成
	if _, err := os.Stat(memoPath); os.IsNotExist(err) {
		if err := c.mode.CreateNewFile(memoPath); err != nil {
			fmt.Println("ファイルの作成中にエラーが発生しました:", err)
			return
		}
	}

	fmt.Printf("メモが作成されました: %s\n", memoPath)

	openInVim(memoPath)
}

// createFromTemplate 関数: テンプレートから新しいメモを作成
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

// createEmptyFile 関数: 空のファイルを作成
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
