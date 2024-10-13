package memo

import (
	"fmt"
	"path/filepath"

	"github.com/banananbo/gomemo/config"
)

type CatMode struct{}

func (m CatMode) DetermineMemoLocation(config *config.Config, category *string) (string, string) {
	if category == nil {
		fmt.Println("「cat」モードに対するカテゴリーが指定されていません")
		return "", ""
	}
	return config.CategoriesDir, *category + ".md"
}

func (m CatMode) CreateNewFile(memoPath string) error {
	templatePath := filepath.Join("templates", "cat_template.md")
	return createFromTemplate(memoPath, templatePath)
}
