package pkg

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
)

// GenerateECDHKeys generates a pair of public and private keys for ECDH.
func GenerateECDHKeys() (publicKey, privateKey []byte, err error) {
	curve := elliptic.P256() // Using P-256 curve

	// Use crypto/ecdsa to generate the keys
	priv, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	public := elliptic.MarshalCompressed(curve, priv.PublicKey.X, priv.PublicKey.Y)
	return public, priv.D.Bytes(), nil
}

// ComputeECDHSecret computes the shared secret using own private key and peer's public key.
func ComputeECDHSecret(privKey, pubKey []byte) ([]byte, error) {
	curve := elliptic.P256() // Using P-256 curve
	x, y := elliptic.UnmarshalCompressed(curve, pubKey)

	if x == nil || y == nil {
		return nil, errors.New("failed to unmarshal public key")
	}

	// Scalar multiplication to compute shared secret
	secretX, _ := curve.ScalarMult(x, y, privKey)
	if secretX == nil {
		return nil, errors.New("failed to compute shared secret")
	}
	return secretX.Bytes(), nil
}