package api

import (
	"crypto/md5"
	"encoding/hex"
)

func calculatePassHash(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}
