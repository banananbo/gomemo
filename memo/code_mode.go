package memo

import (
	"fmt"
	"path/filepath"

	"github.com/banananbo/gomemo/config"
)

type CodeMode struct{}

func (m CodeMode) DetermineMemoLocation(config *config.Config, category *string) (string, string) {
	if category == nil {
		fmt.Println("「code」モードに対するカテゴリーが指定されていません")
		return "", ""
	}
	return config.CodesDir, *category + ".md"
}

func (m CodeMode) CreateNewFile(memoPath string) error {
	templatePath := filepath.Join("templates", "code_template.md")
	return createFromTemplate(memoPath, templatePath)
}
