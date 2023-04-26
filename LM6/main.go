package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Generate an elliptic curve for cryptography
	ellipticCurve := elliptic.P256()
	// Create a new private key based on the chosen elliptic curve
	ecdsaPrivateKey, err := ecdsa.GenerateKey(ellipticCurve, rand.Reader)
	if err != nil {
		panic(err)
	}

	// Derive the public key from the private key
	ecdsaPublicKey := &ecdsaPrivateKey.PublicKey

	// Read input from the user
	fmt.Print("Enter the text to be encrypted: ")
	userInputReader := bufio.NewReader(os.Stdin)
	userInput, err := userInputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// Remove newline character from userInput
	userInput = strings.TrimSpace(userInput)

	// Convert userInput to a byte slice
	inputMessage := []byte(userInput)

	// Generate a random nonce for AES-GCM encryption
	generatedNonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, generatedNonce); err != nil {
		panic(err)
	}

	// Compute a shared secret using the private and public keys
	derivedSecretX, _ := ellipticCurve.ScalarMult(ecdsaPublicKey.X, ecdsaPublicKey.Y, ecdsaPrivateKey.D.Bytes())
	derivedSecret := derivedSecretX.Bytes()

	// Encrypt the inputMessage using AES-GCM
	aesBlock, err := aes.NewCipher(derivedSecret)
	if err != nil {
		panic(err)
	}

	aesGCM, err := cipher.NewGCMWithNonceSize(aesBlock, 12)
	if err != nil {
		panic(err)
	}

	encryptedText := aesGCM.Seal(nil, generatedNonce, inputMessage, nil)

	// Convert the encryptedText to a hexadecimal string for display
	hexadecimalEncryptedText := hex.EncodeToString(encryptedText)

	// Display the encrypted inputMessage to the user
	fmt.Printf("Encrypted text: %s\n", hexadecimalEncryptedText)

	// Decrypt the encryptedText using the same generatedNonce and derivedSecret
	decryptedText, err := aesGCM.Open(nil, generatedNonce, encryptedText, nil)
	if err != nil {
		panic(err)
	}

	// Display the decrypted inputMessage to the user
	fmt.Printf("Decrypted text: %s\n", decryptedText)
}
