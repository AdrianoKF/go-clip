package net

import (
	"net"

	"github.com/AdrianoKF/go-clip/internal/util"
)

type Server struct {
	addr    net.UDPAddr
	Handler HandlerFunc
}

func NewServer(addr net.UDPAddr, handler HandlerFunc) *Server {
	instance := &Server{
		addr,
		handler,
	}
	return instance
}

func (s Server) Listen() {
	conn, err := net.ListenMulticastUDP("udp4", nil, &s.addr)
	if err != nil {
		util.Logger.Error("Error listening on UDP", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, maxDatagramSize)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			util.Logger.Error("Error reading from UDP", err)
			return
		}
		util.Logger.Info("Received UDP packet from ", addr, " with ", n, " bytes")

		if s.Handler != nil {
			go s.Handler(addr, n, buf)
		}
	}
}
