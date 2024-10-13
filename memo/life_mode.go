package memo

import (
	"path/filepath"
	"time"

	"github.com/banananbo/gomemo/config"
)

type LifeMode struct{}

func (m LifeMode) DetermineMemoLocation(config *config.Config, category *string) (string, string) {
	currentTime := time.Now()
	filename := currentTime.Format("200601") + ".md"
	return config.LifeMemoDir, filename
}

func (m LifeMode) CreateNewFile(memoPath string) error {
	templatePath := filepath.Join("templates", "life_template.md")
	return createFromTemplate(memoPath, templatePath)
}
