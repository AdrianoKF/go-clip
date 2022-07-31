package net

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"io"
	"net"

	"github.com/AdrianoKF/go-clip/internal/util"
	"github.com/AdrianoKF/go-clip/pkg/model"
)

type Client struct {
	addr    net.UDPAddr
	handler HandlerFunc
	cipher  cipher.AEAD
}

func NewClient(addr net.UDPAddr, handler HandlerFunc) *Client {
	gcm, err := util.MakeGCMCipher([]byte("secretkey"))
	if err != nil {
		util.Logger.Panic(err)
	}

	instance := &Client{
		addr,
		handler,
		gcm,
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

	nonce := make([]byte, c.cipher.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		util.Logger.Panic(err)
	}

	ciphertext := c.cipher.Seal(nil, nonce, w.Bytes(), nil)

	_, err = conn.Write(append(nonce, ciphertext...))
	if err != nil {
		return err
	}

	return nil
}
