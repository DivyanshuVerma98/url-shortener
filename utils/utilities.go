package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateCode(length int) string {
	randomBytes := make([]byte, length)
	// Read random bytes from the crypto/rand source
	rand.Read(randomBytes)
	// Encode the random bytes to a base64 representation
	randomCode := base64.RawURLEncoding.EncodeToString(randomBytes)
	// Trim the code to the desired length
	if len(randomCode) > length {
		randomCode = randomCode[:length]
	}
	log.Println("Random Bytes", randomBytes)
	log.Println("Generate Code", randomCode)
	return randomCode
}
