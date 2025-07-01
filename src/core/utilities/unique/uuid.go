package unique

import (
	"crypto/rand"
	"fmt"
	"time"
)

func GenerateUUID() string {
	timestamp := time.Now().UnixNano()
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)

	return fmt.Sprintf("%d-%x", timestamp, randomBytes)
}
