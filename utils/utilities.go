package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"strconv"
	"time"
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

func GetCurrentEpoch() int {
	currentTime := time.Now()
	epoch := currentTime.Unix()
	//
	return int(epoch)
}

func ConvertDaysToSec(days int) int {
	return days * 24 * 60 * 60
}

func CreateExpTime(days int) string {
	currEpoch := GetCurrentEpoch()
	buffer := ConvertDaysToSec(days)
	expTime := int64(currEpoch + buffer)
	return strconv.FormatInt(expTime, 10)
}
