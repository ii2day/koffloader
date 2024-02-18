//go:build tools
// +build tools

// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package tools

import (
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/onsi/ginkgo/v2"
	// _ "github.com/gogo/protobuf/gogoproto" // Used for protobuf generation of pkg/k8s/types/slim/k8s
	// _ "golang.org/x/tools/cmd/goimports"
	_ "k8s.io/code-generator"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
