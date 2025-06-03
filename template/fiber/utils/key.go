package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
)

func Hash(msg []byte) []byte {
	hashed := sha256.Sum256(msg)
	return hashed[:]
}

func SignWithECC(key string, hashed []byte) ([]byte, error) {
	// Parse the PEM data.
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		fmt.Println("Failed to decode PEM block")
		return nil, errors.New("failed to decode PEM block")
	}

	if block.Type != "EC PRIVATE KEY" {
		return nil, errors.New("not an EC private key")
	}
	// Parse the DER encoded private key.
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing DER encoded private key:", err)
		return nil, errors.New("failed to decode encoded private key")
	}

	// Sign the hash using the private key.
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashed)
	if err != nil {
		return nil, fmt.Errorf("error signing hash: %v", err)
	}

	// Serialize the signature.
	return asn1.Marshal(struct{ R, S *big.Int }{r, s})
}
