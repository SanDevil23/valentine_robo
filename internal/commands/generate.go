package commands

import "strings"

func IsGenerateCommand(comment string) bool {
	return strings.HasPrefix(comment, "/generate")
}