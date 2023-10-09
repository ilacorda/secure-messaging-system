package main

import (
	"crypto/sha256"
	"fmt"
	"secure-messaging-system/pkg"
)

func main() {
	// Generating ECDH Key Pairs for sender and receiver
	senderPub, senderPriv, err := pkg.GenerateECDHKeys()
	if err != nil {
		fmt.Println("Error generating sender's ECDH keys:", err)
		return
	}

	receiverPub, receiverPriv, err := pkg.GenerateECDHKeys()
	if err != nil {
		fmt.Println("Error generating receiver's ECDH keys:", err)
		return
	}

	// Computing the Shared Secret for both sender and receiver
	senderSharedSecret, err := pkg.ComputeECDHSecret(senderPriv, receiverPub)
	if err != nil {
		fmt.Println("Error computing sender's shared secret:", err)
		return
	}

	receiverSharedSecret, err := pkg.ComputeECDHSecret(receiverPriv, senderPub)
	if err != nil {
		fmt.Println("Error computing receiver's shared secret:", err)
		return
	}

	// Ensure both computed secrets are identical
	// This step might be omitted if confident that the ECDH implementation is correct
	if string(senderSharedSecret) != string(receiverSharedSecret) {
		fmt.Println("Error: Computed shared secrets are not identical!")
		return
	}

	// Processing the Shared Secret to derive the AES encryption key
	key := sha256.Sum256(senderSharedSecret) // Using SHA-256 to produce a 32-byte key for AES-256

	plainText := "This is a test message"
	encryptedText, err := pkg.Encrypt([]byte(plainText), key[:])
	if err != nil {
		fmt.Println("Error encrypting message:", err)
		return
	}

	fmt.Println("Encrypted Text:", string(encryptedText))

	// Use receiver's derived shared secret for decryption
	decryptedText, err := pkg.Decrypt(encryptedText, key[:])
	if err != nil {
		fmt.Println("Error decrypting message:", err)
		return
	}

	fmt.Println("Decrypted Text:", string(decryptedText))
}
