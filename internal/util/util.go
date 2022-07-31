package util

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func InitializeLogging(development bool) {
	var cfg zap.Config
	if development {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	} else {
		cfg = zap.NewProductionConfig()
	}

	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	Logger = l.Sugar()
}

func MakeGCMCipher(key []byte) (cipher.AEAD, error) {
	sha256 := crypto.SHA256.New()
	sha256.Write(key)
	c, err := aes.NewCipher(sha256.Sum(nil))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	return gcm, nil
}
