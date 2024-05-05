
# IST 402 TLS Server, Enigma Machine, and Cryptography Algorithms

This repository contains GoLang implementations for three primary functionalities:
1. **TLS Server and Client** (Module LM3): Secure communication using TLS encryption between a server and a client.
2. **ChaCha20 Encryption/Decryption** (Module LM4): Cryptographic operations using the ChaCha20 cipher.
3. **Elliptic Curve Cryptography (ECC) Example** (Module LM6): Demonstrates encryption and decryption using ECC.
4. **Enigma Machine Simulation** (Module Final): Simulates the historical Enigma machine used for encrypting and decrypting messages.

## Project Structure

This project is divided into several modules:

### LM3 - TLS Server and Client

Implements a basic TLS encrypted communication setup between a server and a client in Go.

#### Key Files
- `main.go`: TLS server implementation.
- `client.go`: TLS client implementation.
- `server.crt` and `server.key`: TLS certificate and key for secure communication.

### LM4 - ChaCha20 Encryption/Decryption

Demonstrates the use of the ChaCha20 encryption algorithm for secure cryptographic operations.

#### Key Files
- `main.go`: Implements the ChaCha20 cipher to encrypt and decrypt messages.

### LM6 - Elliptic Curve Cryptography (ECC)

Provides an example of using elliptic curve cryptography for encryption and decryption.

#### Key Files
- `main.go`: Demonstrates ECC operations including key generation and message encryption/decryption.

### Final - Enigma Machine Simulation

Simulates the Enigma cipher machine, historically known for its role in WWII.

#### Key Files
- `main.go`: Implements the encryption and decryption processes of the Enigma machine using multiple rotors.

## Installation

To run these projects, ensure Go is installed on your machine. Follow these steps:

1. **Clone the repository:**
   ```bash
   git clone git@github.com:jrobin11/IST-402.git
   cd IST-402
   ```

2. **Install Go dependencies for each module:**
   Navigate to each module directory and run:
   ```bash
   go mod tidy
   ```

## Usage

### LM3 - TLS Server and Client

- **Server:**
  ```bash
  cd LM3
  go run main.go
  ```

- **Client:**
  ```bash
  go run client.go
  ```

### LM4 - ChaCha20 Encryption/Decryption

- Run the ChaCha20 example:
  ```bash
  cd LM4
  go run main.go
  ```

### LM6 - ECC Example

- Run the ECC example:
  ```bash
  cd LM6
  go run main.go
  ```

### Final - Enigma Machine Simulation

- Run the Enigma machine simulation:
  ```bash
  cd Final
  go run main.go
  ```

## Contributing

Contributions are welcome. Please ensure to update tests as appropriate.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

- This project was created as part of the IST 402 class at Penn State University.

