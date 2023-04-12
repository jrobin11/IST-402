# Assignment Client GoLang TLS Programming Copy
Using GoLang perform TLS encryption and decryption on a user string input.

Step 1: Make sure to create the following files in the project working for example C:/Users/joeoa/GolandProjects/TLS_Server

// Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048

// Key considerations for algorithm "ECDSA" (X25519 || ≥ secp384r1)
// https://safecurves.cr.yp.to/
//List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key

// Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)

openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

<img width="587" alt="Screenshot 2023-04-11 at 9 17 34 PM" src="https://user-images.githubusercontent.com/73866458/231322416-f9e591ce-3f72-42d8-a0c7-9bf81b00a91a.png">
