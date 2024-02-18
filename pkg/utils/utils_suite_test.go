// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package utils_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Utils Suite")
}
