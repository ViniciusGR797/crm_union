package security

import (
	"crypto/sha256"
	"fmt"
)

func SHA256Encode(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}
