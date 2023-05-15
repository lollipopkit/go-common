package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func Contains[T string | int | int32 | int64 | float64](s []T, e T) bool {
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
