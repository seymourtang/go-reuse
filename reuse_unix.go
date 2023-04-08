//go:build linux || darwin || dragonfly || freebsd || netbsd || openbsd

package reuse

import (
	"syscall"

	"golang.org/x/sys/unix"
)

func control(network, address string, conn syscall.RawConn) error {
	return conn.Control(func(fd uintptr) {
		if err := unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1); err != nil {
			panic(err)
		}

		if err := unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1); err != nil {
			panic(err)
		}
	})
}
