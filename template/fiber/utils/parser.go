package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func ParseEccPublicKeyFromPEM(key string) (ecdsaPubKey *ecdsa.PublicKey, err error) {
	// Parse the PEM data.
	block, _ := pem.Decode([]byte(key))
	if block == nil || block.Type != "PUBLIC KEY" {
		fmt.Println("Failed to decode PEM block containing public key")
		return
	}

	// Parse the DER encoded public key.
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing DER encoded public key:", err)
		return
	}

	// Convert the parsed public key to an ecdsa.PublicKey object.
	ecdsaPubKey, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Failed to convert to ECDSA public key")
		return
	}

	return
}
