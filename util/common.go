package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func Contains[T comparable](s []T, e T) bool {
	for idx := range s {
		if s[idx] == e {
			return true
		}
	}
	return false
}

// Get SHA256
func GetSHA256(r io.Reader) string {
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}
