package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

const (
	errCipherInit      = "failed to initialize AES cipher: %w"
	errShortCiphertext = "ciphertext is too short"
	errGeneratingIV    = "failed to generate initialization vector: %w"
)

// Encrypt takes a plaintext and key and returns the encrypted text.
func Encrypt(plainText, key []byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 { // AES-128, AES-192, AES-256
		return nil, fmt.Errorf("invalid key size: %d", len(key))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf(errCipherInit, err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf(errGeneratingIV, err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)
	return ciphertext, nil
}

// Decrypt takes an encrypted text and key and returns the original text.
func Decrypt(cipherText, key []byte) ([]byte, error) {
	if len(cipherText) < aes.BlockSize {
		return nil, errors.New(errShortCiphertext)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf(errCipherInit, err)
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return cipherText, nil
}
