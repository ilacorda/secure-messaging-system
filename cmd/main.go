package main

import (
	"crypto/sha256"
	"github.com/rs/zerolog/log"
	"secure-messaging-system/pkg"
)

func main() {
	// Initialize logger
	pkg.Setup()

	// Generating ECDH Key Pairs for sender and receiver
	senderPub, senderPriv, err := pkg.GenerateECDHKeys()
	if err != nil {
		log.Error().Err(err).Msg("Error generating sender's ECDH keys")
		return
	}

	receiverPub, receiverPriv, err := pkg.GenerateECDHKeys()
	if err != nil {
		log.Error().Err(err).Msg("Error generating receiver's ECDH keys")
		return
	}

	// Computing the Shared Secret for both sender and receiver
	senderSharedSecret, err := pkg.ComputeECDHSecret(senderPriv, receiverPub)
	if err != nil {
		log.Error().Err(err).Msg("Error computing sender's shared secret")
		return
	}

	receiverSharedSecret, err := pkg.ComputeECDHSecret(receiverPriv, senderPub)
	if err != nil {
		log.Error().Err(err).Msg("Error computing receiver's shared secret")
		return
	}

	// Ensure both computed secrets are identical
	if string(senderSharedSecret) != string(receiverSharedSecret) {
		log.Error().Msg("Computed shared secrets are not identical")
		return
	}

	// Processing the Shared Secret to derive the AES encryption key
	key := sha256.Sum256(senderSharedSecret) // Using SHA-256 to produce a 32-byte key for AES-256

	plainText := "This is a test message"
	encryptedText, err := pkg.Encrypt([]byte(plainText), key[:])
	if err != nil {
		log.Error().Err(err).Msg("Error encrypting message")
		return
	}

	log.Info().Str("Encrypted Text", string(encryptedText)).Msg("Message encrypted successfully")

	// Use receiver's derived shared secret for decryption
	decryptedText, err := pkg.Decrypt(encryptedText, key[:])
	if err != nil {
		log.Error().Err(err).Msg("Error decrypting message")
		return
	}

	log.Info().Str("Decrypted Text", string(decryptedText)).Msg("Message decrypted successfully")
}
