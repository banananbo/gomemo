package config

import (
	"encoding/json"
	"os"
)

// Config 構造体は設定ファイルの内容を表す
type Config struct {
	DefaultMemoDir string `json:"defaultMemoDir"`
	LifeMemoDir    string `json:"lifeMemoDir"`
}

// LoadConfig は指定されたファイルパスから設定を読み込む
func LoadConfig(filepath string) (*Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
