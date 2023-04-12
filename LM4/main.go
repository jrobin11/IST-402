package main

import (
	"crypto/rand"  // package for generating random bytes
	"encoding/hex" // package for encoding and decoding hexadecimal strings
	"fmt"
	"os"

	"golang.org/x/crypto/chacha20poly1305" // package for the ChaCha20-Poly1305 encryption algorithm
)

func main() {
	// Get user input
	fmt.Print("Enter a message to encrypt: ") // read user input from stdin
	message, err := readInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // exit program if there's an error
	}

	// Generate a random 256-bit key
	key := make([]byte, 32)                   // create a byte slice of length 12 (the length of a nonce for ChaCha20-Poly1305)
	if _, err := rand.Read(key); err != nil { // fill the byte slice with random bytes
		fmt.Println(err)
		os.Exit(1)
	}

	// Generate a random nonce
	nonce := make([]byte, chacha20poly1305.NonceSize) // create a byte slice of length 12 (the length of a nonce for ChaCha20-Poly1305)
	if _, err := rand.Read(nonce); err != nil {       // fill the byte slice with random bytes
		fmt.Println(err)
		os.Exit(1)
	}

	// Encrypt the message using ChaCha20-Poly1305
	ciphertext := encrypt(message, key, nonce)

	// Print the encrypted message and key
	fmt.Printf("Encrypted message: %s\n", hex.EncodeToString(ciphertext)) // encode the ciphertext as a hexadecimal string for printing
	fmt.Printf("Key: %s\n", hex.EncodeToString(key))                      // encode the key as a hexadecimal string for printing
	fmt.Printf("Nonce: %s\n", hex.EncodeToString(nonce))                  // encode the nonce as a hexadecimal string for printing

	// Decrypt the message using ChaCha20-Poly1305
	plaintext := decrypt(ciphertext, key, nonce) // decrypt the ciphertext using the generated key and nonce

	// Print the decrypted message
	fmt.Printf("Decrypted message: %s\n", plaintext)
}

func readInput() ([]byte, error) {
	input := make([]byte, 1024)    // create a byte slice of length 1024 to hold user input
	n, err := os.Stdin.Read(input) // read user input from stdin
	if err != nil {
		return nil, err // return an error if there's a problem reading user input
	}
	return input[:n-1], nil // return the user input as a byte slice (without the newline character)
}

func encrypt(plaintext, key, nonce []byte) []byte {
	// Create a new ChaCha20-Poly1305 cipher
	block, err := chacha20poly1305.New(key) // create a new ChaCha20-Poly1305 cipher using the generated key
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Encrypt the plaintext
	ciphertext := make([]byte, 0, len(plaintext)+block.Overhead()) // create a byte slice to hold the encrypted ciphertext
	ciphertext = block.Seal(ciphertext, nonce, plaintext, nil)     // encrypt the plaintext using the generated nonce and the ChaCha20-Poly1305 cipher

	return ciphertext // return the encrypted ciphertext as a byte slice
}

// decrypt takes in a ciphertext byte slice, a key byte slice, and a nonce byte slice
// and returns a plaintext byte slice. It uses the ChaCha20-Poly1305 cipher to decrypt
// the ciphertext using the provided key and nonce.
func decrypt(ciphertext, key, nonce []byte) []byte {
	// Create a new ChaCha20-Poly1305 cipher with the provided key
	block, err := chacha20poly1305.New(key)
	if err != nil {
		// If there is an error creating the cipher, print the error and exit the program
		fmt.Println(err)
		os.Exit(1)
	}

	// Decrypt the ciphertext using the provided nonce and the ChaCha20-Poly1305 cipher
	plaintext, err := block.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		// If there is an error decrypting the ciphertext, print the error and exit the program
		fmt.Println(err)
		os.Exit(1)
	}

	// Return the plaintext
	return plaintext
}
