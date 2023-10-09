package pkg

import (
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"os"
)

const envEncryptionKey = "TEST_ENCRYPTION_KEY"

// GetEncryptionKey fetches and validates the encryption key from environment variables.
func GetEncryptionKey() ([]byte, error) {
	key := os.Getenv(envEncryptionKey)
	if len(key) == 0 {
		return nil, fmt.Errorf("environment variable %s is not set", envEncryptionKey)
	}

	if len(key) != 16 && len(key) != 24 && len(key) != 32 { // AES-128, AES-192, AES-256
		return nil, fmt.Errorf("invalid key size: %d", len(key))
	}
	return []byte(key), nil
}

// GenerateECDHKeys generates a pair of public and private keys for ECDH.
func GenerateECDHKeys() (publicKey, privateKey []byte, err error) {
	curve := elliptic.P256() // Using P-256 curve
	private, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	public := elliptic.MarshalCompressed(curve, x, y)
	return public, private, nil
}

// ComputeECDHSecret computes the shared secret using own private key and peer's public key.
func ComputeECDHSecret(privKey, pubKey []byte) ([]byte, error) {
	curve := elliptic.P256() // Using P-256 curve
	x, y := elliptic.UnmarshalCompressed(curve, pubKey)

	if !curve.IsOnCurve(x, y) {
		return nil, errors.New("invalid elliptic curve point")
	}

	// Compute the secret using elliptic curve scalar multiplication
	x, _ = curve.ScalarMult(x, y, privKey)
	return x.Bytes(), nil
}
