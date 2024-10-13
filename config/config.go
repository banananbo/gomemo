package config

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"path/filepath"
)

// Config 構造体は設定ファイルの内容を表す
type Config struct {
	RootDir        string `json:"rootDir"`
	DefaultMemoDir string `json:"defaultMemoDir"`
	LifeMemoDir    string `json:"lifeMemoDir"`
	CategoriesDir  string `json:"CategoriesDir"`
	CodesDir       string `json:"CodesDir"`
}

//go:embed config.json
var embeddedConfig []byte

func LoadConfig() (*Config, error) {
	var config Config
	if err := json.Unmarshal(embeddedConfig, &config); err != nil {
		return nil, err
	}
	// RootDir を含めてディレクトリのパスを更新
	config.DefaultMemoDir = filepath.Join(config.RootDir, config.DefaultMemoDir)
	config.LifeMemoDir = filepath.Join(config.RootDir, config.LifeMemoDir)
	config.CategoriesDir = filepath.Join(config.RootDir, config.CategoriesDir)
	config.CodesDir = filepath.Join(config.RootDir, config.CodesDir)
	return &config, nil
}

func main() {
	config, err := LoadConfig()

	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	fmt.Println("DefaultMemoDir:", config.DefaultMemoDir)
	fmt.Println("LifeMemoDir:", config.LifeMemoDir)
}
