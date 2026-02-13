package commands

import "strings"

func isGenerateCommand(comment string) bool {
	return strings.HasPrefix(comment, "/generate")
}