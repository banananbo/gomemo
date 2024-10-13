package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/banananbo/gomemo/config"
	"github.com/banananbo/gomemo/memo"
)

func main() {
	// Configのロード
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Configのロードに失敗しました:", err)
		return
	}

	// MemoContextを初期化
	memoContext := memo.MemoContext{}

	// モードに応じて戦略をセットし、メモを開く
	if len(os.Args) < 2 {
		// 引数がなければデフォルトモード
		memoContext.SetMode(memo.DefaultMode{})
		memoContext.OpenMemo(config, nil)
	} else {
		command := os.Args[1]
		switch {
		case command == "life":
			memoContext.SetMode(memo.LifeMode{})
			memoContext.OpenMemo(config, nil)

		case strings.HasPrefix(command, "cat="):
			category := strings.TrimPrefix(command, "cat=")
			memoContext.SetMode(memo.CatMode{})
			memoContext.OpenMemo(config, &category)

		case strings.HasPrefix(command, "code="):
			category := strings.TrimPrefix(command, "code=")
			memoContext.SetMode(memo.CodeMode{})
			memoContext.OpenMemo(config, &category)

		default:
			memoContext.SetMode(memo.DefaultMode{})
			memoContext.OpenMemo(config, nil)
		}
	}
}
