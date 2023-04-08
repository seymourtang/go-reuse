//go:build !windows && !linux && !darwin && !dragonfly && !freebsd && !netbsd && !openbsd

package reuse

import "syscall"

func control(network, address string, c syscall.RawConn) error {
	return nil
}
