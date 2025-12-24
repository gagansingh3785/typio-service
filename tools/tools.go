//go:build tools
// +build tools

package tools

import (
	_ "mvdan.cc/gofumpt"

	_ "github.com/daixiang0/gci/cmd/gci"
	_ "github.com/golangci/golangci-lint/v2/cmd/golangci-lint"
)
