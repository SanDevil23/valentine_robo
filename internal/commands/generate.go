package commands
// Copyright (c) 2026 SanDevil23
// SPDX-License-Identifier: Apache-2.0

import "strings"

func IsGenerateCommand(comment string) bool {
	return strings.HasPrefix(comment, "/generate")
}