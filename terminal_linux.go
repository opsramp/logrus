// Based on ssh/terminal:
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !appengine,!gopherjs

package logrus

import "golang.org/x/sys/unix"

import "syscall"

const ioctlReadTermios = unix.TCGETS

type Termios unix.Termios

func GetCurrentThreadId() int {
	return syscall.Gettid()
}
