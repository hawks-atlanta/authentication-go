package random

import "crypto/rand"

func Bytes(length int) (buf []byte) {
	buf = make([]byte, length)
	rand.Read(buf)
	return buf
}
