package llm

import (
	"errors"
	"strings"
)

var blockedPaths = []string{
	".github/workflows",
	".env",
	"infra/",
	"helm/",
}

const maxFiles = 10
const maxFileSize = 10000 // 10 KB

func ValidateFiles(files []File) error {

	if len(files) > maxFiles {
		return errors.New("too many files generated")
	}

	for _, f := range files {

		if f.Path == "" {
			return errors.New("empty file path")
		}

		for _, blocked := range blockedPaths {
			if strings.Contains(f.Path, blocked) {
				return errors.New("blocked path detected: " + f.Path)
			}
		}

		if len(f.Content) > maxFileSize {
			return errors.New("file too large: " + f.Path)
		}
	}

	return nil
}
