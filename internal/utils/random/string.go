package random

import (
	"crypto/rand"
	"encoding/hex"
)

func String(length int) string {
	var buf = make([]byte, length)
	rand.Read(buf)
	return hex.EncodeToString(buf)[:length]
}
