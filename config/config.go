package config

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

// Config 構造体は設定ファイルの内容を表す
type Config struct {
	DefaultMemoDir string `json:"defaultMemoDir"`
	LifeMemoDir    string `json:"lifeMemoDir"`
}

//go:embed config.json
var embeddedConfig []byte

func LoadConfig() (*Config, error) {
	var config Config
	if err := json.Unmarshal(embeddedConfig, &config); err != nil {
		return nil, err
	}
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

// LoadConfig は指定されたファイルパスから設定を読み込む
// func LoadConfig(filepath string) (*Config, error) {
// 	file, err := os.Open(filepath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var config Config
// 	if err := json.NewDecoder(file).Decode(&config); err != nil {
// 		return nil, err
// 	}
// 	return &config, nil
// }
