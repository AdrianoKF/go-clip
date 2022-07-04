package net

import (
	"bytes"
	"encoding/json"
	"net"

	"github.com/AdrianoKF/go-clip/pkg/model"
)

type Client struct {
	addr    net.UDPAddr
	handler HandlerFunc
}

func NewClient(addr net.UDPAddr, handler HandlerFunc) *Client {
	instance := &Client{
		addr,
		handler,
	}
	return instance
}

func (c Client) SendEvent(msg model.ClipboardUpdated) error {
	conn, err := net.DialUDP("udp4", nil, &c.addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	buf := make([]byte, 0)
	w := bytes.NewBuffer(buf)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)

	_, err = conn.Write(w.Bytes())
	if err != nil {
		return err
	}

	return nil
}
