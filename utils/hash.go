package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func HmacSha256Hash(data []byte, key string) string {
	hmacSha256Hash := hmac.New(sha256.New, []byte(key))
	hmacSha256Hash.Write(data)
	return base64.StdEncoding.EncodeToString([]byte(hmacSha256Hash.Sum(nil)))
}

func Sha256(data string) string {
	sumed := sha256.Sum256([]byte(data))
	return hex.EncodeToString(sumed[:])
}

func Sha1(data string) string {
	hasher := sha1.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
