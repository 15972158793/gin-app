package utils

import (
	"crypto/md5"
	"encoding/hex"
)

var secret = "pdf20221211"

func MD5(value string) string {
	m := md5.New()
	m.Write([]byte(secret))
	return hex.EncodeToString(m.Sum([]byte(value)))
}
