package reuse

import (
	"context"
	"net"
)

var listenConfig = net.ListenConfig{
	Control: control,
}

func Listen(network, address string) (net.Listener, error) {
	return listenConfig.Listen(context.Background(), network, address)
}

func ListenPacket(network, address string) (net.PacketConn, error) {
	return listenConfig.ListenPacket(context.Background(), network, address)
}
