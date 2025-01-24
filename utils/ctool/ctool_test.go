package ctool

import (
	"fmt"
	"testing"
)

func TestCryptoWithSalt(t *testing.T) {
	plainTests := []string{
		"abc",
		"hello world",
		"hello world",
		"hello world!",
	}
	cipherText := make([]string, 0)
	for _, p := range plainTests {
		cipherText = append(cipherText, CryptoWithSalt(p))
	}

	for _, s := range cipherText {
		fmt.Println(s)
	}
}
