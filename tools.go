//go:build tools
// +build tools

package clickhouse_go_rows_utils

// This file directly fixes tool dependencies in your go.mod
// with initial tags. It's needed to have same dependencies
// between engineers of the project.
// You can bump version manually with 'go get package@version'

import (
	_ "go.uber.org/mock/mockgen"
)
