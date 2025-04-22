package utils

import (
	"crypto/rand"
	"fmt"
	"time"
)

func GenerateCustomGUID() string {
	timestamp := time.Now().UnixNano()
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)

	return fmt.Sprintf("%d-%x", timestamp, randomBytes)
}
