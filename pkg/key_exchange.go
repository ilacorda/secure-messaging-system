package pkg

import (
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/rs/zerolog/log"
)

// GenerateECDHKeys generates a pair of public and private keys for ECDH.
func GenerateECDHKeys() (publicKey, privateKey []byte, err error) {
	curve := elliptic.P256() // Using P-256 curve
	log.Info().Str("curve", "P-256").Msg("Generating ECDH keys using curve")

	private, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate ECDH keys")
		return nil, nil, err
	}

	public := elliptic.MarshalCompressed(curve, x, y)
	log.Debug().Hex("publicKey", public).Msg("Generated public key for ECDH")
	return public, private, nil
}

// ComputeECDHSecret computes the shared secret using own private key and peer's public key.
func ComputeECDHSecret(privKey, pubKey []byte) ([]byte, error) {
	curve := elliptic.P256() // Using P-256 curve
	log.Info().Str("curve", "P-256").Msg("Computing ECDH secret using curve")

	x, y := elliptic.UnmarshalCompressed(curve, pubKey)
	if !curve.IsOnCurve(x, y) {
		log.Warn().Str("point", "elliptic curve").Msg("Invalid point on curve detected")
		return nil, errors.New("invalid elliptic curve point")
	}

	// Compute the secret using elliptic curve scalar multiplication
	x, _ = curve.ScalarMult(x, y, privKey)
	sharedSecret := x.Bytes()
	log.Debug().Msg("Successfully computed ECDH secret")
	return sharedSecret, nil
}
