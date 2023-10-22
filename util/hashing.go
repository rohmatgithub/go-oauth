package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/cespare/xxhash/v2"
	"strconv"
)

func CheckSumWithXXHASH(content []byte) (checksum string) {
	hash := xxhash.Sum64(content)
	return strconv.Itoa(int(hash))
}

func CheckSumWithMD5(content []byte) (checksum string) {
	hash := md5.New()
	hash.Write(content)
	hashInBytes := hash.Sum(nil)[:16]
	return hex.EncodeToString(hashInBytes)
}

func CheckSumWithSha256(content []byte) string {
	result := sha256.Sum256(content)
	return hex.EncodeToString(result[:])
}

func CheckSumWithSha512(content []byte) string {
	result := sha512.Sum512(content)
	return hex.EncodeToString(result[:])
}

func GetHmacSHA512(input string, key string) string {
	mac := hmac.New(sha512.New, []byte(key))
	mac.Write([]byte(input))
	return hex.EncodeToString(mac.Sum(nil))
}

func HashingPassword(password string, salt string) string {
	return CheckSumWithSha512([]byte(password + salt))
}

func CheckIsPasswordMatch(passwordInput string, passwordDB string, salt string) bool {
	return HashingPassword(passwordInput, salt) == passwordDB
}
