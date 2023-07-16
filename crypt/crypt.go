package crypt

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func GetSHA256(r io.Reader) string {
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}