package net

import "net"

const (
	maxDatagramSize = 8192
)

type HandlerFunc func(*net.UDPAddr, int, []byte)
