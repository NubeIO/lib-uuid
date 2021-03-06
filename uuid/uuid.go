package uuid

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"
)

func UUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x%x%x", uuid[0:4], uuid[4:6], uuid[6:8]), nil
}

func ShortUUID(prefix ...string) (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	if len(prefix) > 0 {
		uuid = fmt.Sprintf("%s_%X%X", prefix[0], b[0:4], b[4:6])
	} else {
		uuid = fmt.Sprintf("id_%X%X", b[0:4], b[4:6])
	}
	return
}

func UniqueNumber() (uuid int64) {
	return time.Now().UnixNano() / (1 << 22)
}
