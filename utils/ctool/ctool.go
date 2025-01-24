package ctool

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

const salt = "0x00447953686f70"

func CryptoWithSalt(plaintext string) string {
	return fmt.Sprintf(
		"%x",
		md5.Sum([]byte(
			fmt.Sprintf("%x%x",
				sha256.Sum256([]byte(plaintext)),
				sha256.Sum256([]byte(salt)),
			),
		)),
	)
}
