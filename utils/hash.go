package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

type Hash struct {
	aesGCM cipher.AEAD
}

func (h *Hash) EncryptAES(data string) (string, error) {
	byteData := []byte(data)

	nonce := make([]byte, h.aesGCM.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	cipherText := h.aesGCM.Seal(nonce, nonce, byteData, nil)
	return string(cipherText), nil
}

func (h *Hash) DecryptAES(data string) (string, error) {
	byteData := []byte(data)

	nonceSize := h.aesGCM.NonceSize()
	if len(byteData) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, cipherText := byteData[:nonceSize], byteData[nonceSize:]

	plainText, err := h.aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func NewHash(key string) (*Hash, error) {
	if key == "" {
		return &Hash{}, nil
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &Hash{
		aesGCM: aesGCM,
	}, nil
}
