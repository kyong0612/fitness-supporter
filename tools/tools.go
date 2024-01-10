//go:build tools
// +build tools

package tools

import (
	_ "connectrpc.com/connect/cmd/protoc-gen-connect-go"
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/cosmtrek/air"
	_ "github.com/fullstorydev/grpcurl/cmd/grpcurl"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "golang.org/x/vuln/cmd/govulncheck"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
