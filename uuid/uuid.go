package uuid

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"
)

func UUID(prefix ...string) string {
	u := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, u)
	if n != len(u) || err != nil {
		return "---error-uuid---"
	}
	u[8] = u[8]&^0xc0 | 0x80
	u[6] = u[6]&^0xf0 | 0x40

	uuid := fmt.Sprintf("%x%x%x", u[0:4], u[4:6], u[6:8])
	if len(prefix) > 0 {
		uuid = fmt.Sprintf("%s_%s", prefix[0], uuid)
	}
	return uuid
}

func ShortUUID(prefix ...string) string {
	u := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, u)
	if n != len(u) || err != nil {
		return "-error-uuid-"
	}
	uuid := fmt.Sprintf("%x%x", u[0:4], u[4:6])
	if len(prefix) > 0 {
		uuid = fmt.Sprintf("%s_%s", prefix[0], uuid)
	}
	return uuid
}

func UniqueNumber() (uuid int64) {
	return time.Now().UnixNano() / (1 << 22)
}
