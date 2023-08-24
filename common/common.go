package common

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	"github.com/google/uuid"
)

func GetUUID() string {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	return id.String()
}

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
