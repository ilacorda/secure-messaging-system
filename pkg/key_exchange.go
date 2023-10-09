package pkg

import (
	"crypto/elliptic"
	"crypto/rand"
	"errors"
)

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
