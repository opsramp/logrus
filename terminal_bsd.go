// +build freebsd openbsd netbsd dragonfly
// +build !appengine,!gopherjs

package logrus

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TIOCGETA

type Termios unix.Termios

func GetCurrentThreadId() int {
	//Since FreeBSD does not have a systemcall for thread id as of now. We are using process id instead
	return unix.Getpid()
}

