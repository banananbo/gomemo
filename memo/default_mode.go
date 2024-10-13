package memo

import (
	"time"

	"github.com/banananbo/gomemo/config"
)

type DefaultMode struct{}

func (m DefaultMode) DetermineMemoLocation(config *config.Config, category *string) (string, string) {
	currentTime := time.Now()
	filename := currentTime.Format("20060102") + ".md"
	return config.DefaultMemoDir, filename
}

func (m DefaultMode) CreateNewFile(memoPath string) error {
	return createEmptyFile(memoPath)
}
