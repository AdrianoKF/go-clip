package net

import (
	"crypto/cipher"
	"net"

	"github.com/AdrianoKF/go-clip/internal/util"
)

type Server struct {
	addr    net.UDPAddr
	Handler HandlerFunc
	cipher  cipher.AEAD
}

func NewServer(addr net.UDPAddr, encKey string, handler HandlerFunc) *Server {
	var cipher cipher.AEAD = nil

	if encKey != "" {
		var err error
		cipher, err = util.MakeGCMCipher([]byte(encKey))
		if err != nil {
			util.Logger.Panic(err)
		}
	} else {
		util.Logger.Warn("Using unencrypted connection")
	}

	instance := &Server{
		addr,
		handler,
		cipher,
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
		util.Logger.Debug("Received UDP packet from ", addr, " with ", n, " bytes")

		if s.Handler != nil {
			var plaintext []byte
			if s.cipher != nil {
				var err error
				nonceLen := s.cipher.NonceSize()
				plaintext, err = s.cipher.Open(nil, buf[:nonceLen], buf[nonceLen:n], nil)
				if err != nil {
					util.Logger.Error(err)
					continue
				}
			} else {
				plaintext = buf[:n]
			}

			util.Logger.Debug("Decrypted event data:", string(plaintext))

			go s.Handler(addr, n, plaintext)
		}
	}
}
