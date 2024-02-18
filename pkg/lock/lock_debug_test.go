// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of koffloader-io

//go:build lockdebug
// +build lockdebug

package lock_test

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/koffloader-io/koffloader/pkg/lock"
)

var _ = Describe("LockFast", Label("unitest"), func() {

	// it is daemon , add more test here
	It("test debug lock", func() {

		l := &lock.Mutex{}

		l.Lock()
		l.Unlock()
	})
	It("test debug Rlock", func() {
		l := &lock.RWMutex{}
		l.RLock()
		l.RUnlock()
		l.Lock()
		l.Unlock()
	})
})
