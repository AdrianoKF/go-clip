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

func NewClient(addr net.UDPAddr, encKey string, handler HandlerFunc) *Client {
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

	instance := &Client{
		addr,
		handler,
		cipher,
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

	if c.cipher != nil {
		nonce := make([]byte, c.cipher.NonceSize())
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			util.Logger.Panic(err)
		}

		ciphertext := c.cipher.Seal(nil, nonce, w.Bytes(), nil)

		_, err = conn.Write(append(nonce, ciphertext...))
		if err != nil {
			return err
		}
	} else {
		_, err = conn.Write(w.Bytes())
		if err != nil {
			return err
		}
	}

	return nil
}
