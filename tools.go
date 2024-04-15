//go:build tools
// +build tools

// This file is used to track development dependencies.
// This ensures `go mod` can detect and include them in the project's dependencies.

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/lib/pq"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
	_ "github.com/vektra/mockery/v2"
)
