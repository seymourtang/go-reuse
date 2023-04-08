//go:build windows

package reuse

import (
	"syscall"

	"golang.org/x/sys/windows"
)

func control(network, address string, conn syscall.RawConn) error {
	return conn.Control(func(fd uintptr) {
		if err := windows.SetsockoptInt(windows.Handle(fd), windows.SOL_SOCKET, windows.SO_REUSEADDR, 1); err != nil {
			panic(err)
		}
	})
}
