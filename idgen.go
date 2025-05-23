package idgen

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
)

var bufPool = sync.Pool{
	New: func() any {
		b := make([]byte, 16)
		return &b
	},
}

// Generate возвращает URL-safe base64 строку длиной ~ceil(byteLength * 4 / 3) символов.
// Подходит для использования в ссылках, токенах и идентификаторах.
func Generate(byteLength int) (string, error) {
	bufPtr := bufPool.Get().(*[]byte)
	defer bufPool.Put(bufPtr)

	buf := *bufPtr
	if cap(buf) < byteLength {
		buf = make([]byte, byteLength)
	}
	buf = buf[:byteLength]

	_, err := rand.Read(buf)
	if err != nil {
		return "", fmt.Errorf("idgen: failed to generate random bytes: %w", err)
	}

	return base64.RawURLEncoding.EncodeToString(buf), nil
}
