package main

import (
	"bufio"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create an elliptic curve instance
	ec := elliptic.P256()

	// Generate a private key for the curve
	privKey, pubX, pubY, err := elliptic.GenerateKey(ec, rand.Reader)
	if err != nil {
		panic(err)
	}

	// Prompt user for a string to encrypt
	fmt.Print("Enter a string to encrypt: ")
	reader := bufio.NewReader(os.Stdin)
	inputStr, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// Remove newline character from input
	inputStr = strings.TrimSuffix(inputStr, "\n")

	// Convert the input string to bytes
	msgBytes := []byte(inputStr)

	// Generate a random scalar value k
	k, err := rand.Int(rand.Reader, ec.Params().N)
	if err != nil {
		panic(err)
	}

	// Compute the ephemeral public key R = k * G (G is the curve's generator point)
	Rx, Ry := ec.ScalarBaseMult(k.Bytes())
	R := elliptic.Marshal(ec, Rx, Ry)

	// Compute the shared secret S = k * (pubX, pubY)
	Sx, Sy := ec.ScalarMult(pubX, pubY, k.Bytes())
	S := elliptic.Marshal(ec, Sx, Sy)

	// Compute the hash of the shared secret S to obtain the symmetric key K
	hashedS := Hash(S)
	K := hashedS[:16] // Use the first 16 bytes as the key for simplicity

	// Encrypt the message using XOR with the symmetric key K
	encMsg := make([]byte, len(msgBytes))
	for i := range msgBytes {
		encMsg[i] = msgBytes[i] ^ K[i%len(K)]
	}

	// Convert the encrypted message and ephemeral public key R to hexadecimal strings
	hexEncMsg := hex.EncodeToString(encMsg)
	hexR := hex.EncodeToString(R)

	// Output the encrypted message and ephemeral public key R
	fmt.Printf("Encrypted [message]: %s\n", hexEncMsg)
	fmt.Printf("Ephemeral [public key]: %s\n", hexR)

	// Extract the ephemeral public key R and compute the shared secret S using the private key
	R, _ = hex.DecodeString(hexR)
	Rx, Ry = elliptic.Unmarshal(ec, R)
	Sx, Sy = ec.ScalarMult(Rx, Ry, privKey)

	// Compute the hash of the shared secret S to obtain the symmetric key K
	S = elliptic.Marshal(ec, Sx, Sy)
	hashedS = Hash(S)
	K = hashedS[:16] // Use the first 16 bytes as the key for simplicity

	// Decrypt the encrypted message using XOR with the symmetric key K
	encMsg, _ = hex.DecodeString(hexEncMsg)
	decMsg := make([]byte, len(encMsg))
	for i := range encMsg {
		decMsg[i] = encMsg[i] ^ K[i%len(K)]
	}

	// Output the decrypted message
	fmt.Printf("Decrypted [message]: %s\n", decMsg)
}

// Hash computes the SHA-256 hash of a byte slice
func Hash(input []byte) []byte {
	hashed := sha256.Sum256(input)
	return hashed[:]
}
