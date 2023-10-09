package test

import (
	"secure-messaging-system/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testMsg       = "This is a test message"
	encryptErrMsg = "Encrypt should return an error"
)

func Test_EncryptDecryptLogic(t *testing.T) {
	testCases := []struct {
		name       string
		plainText  []byte
		key        []byte
		shouldFail bool
	}{
		{
			name:       "Valid encryption and decryption",
			plainText:  []byte(testMsg),
			key:        []byte("ThisKeyIsExactly32BytesLong1234!"),
			shouldFail: false,
		},
		{
			name:       "Invalid key size",
			plainText:  []byte(testMsg),
			key:        []byte("InvalidKey"),
			shouldFail: true,
		},
		{
			name:       "Short ciphertext",
			plainText:  []byte(testMsg),
			key:        []byte("TestEncryptionKey123"),
			shouldFail: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ciphertext, err := pkg.Encrypt(tc.plainText, tc.key)

			if tc.shouldFail {
				assert.Error(t, err, encryptErrMsg)
				assert.Nil(t, ciphertext, "Ciphertext should be nil")
			} else {
				assert.NoError(t, err, encryptErrMsg)

				decryptedText, err := pkg.Decrypt(ciphertext, tc.key)
				assert.NoError(t, err, "Decrypt should not return an error")
				assert.Equal(t, tc.plainText, decryptedText, "Decrypted text should match the original plaintext")
			}
		})
	}
}
