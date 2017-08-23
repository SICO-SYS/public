/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package public

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func EncryptWithSha256(v string) string {
	hash := sha256.New()
	hash.Write([]byte(v))
	return hex.EncodeToString(hash.Sum(nil))
}

func EncryptBytesWithSha256(value []byte) []byte {
	hash := sha256.New()
	hash.Write(value)
	return hash.Sum(nil)
}

func EncryptWithHmacSha256(k string, v string) string {
	s := hmac.New(sha256.New, []byte(k))
	s.Write([]byte(v))
	return hex.EncodeToString(s.Sum(nil))
}

func EncryptBytesWithHmacSha256(key []byte, value string) []byte {
	s := hmac.New(sha256.New, key)
	s.Write([]byte(value))
	return s.Sum(nil)
}

func Hmac256ToBase64(key string, str string, IsUrl bool) string {
	s := hmac.New(sha256.New, []byte(key))
	s.Write([]byte(str))
	return EncodingBase64(s.Sum(nil), IsUrl)
}

func EncryptWithSha1(v string) string {
	hash := sha1.New()
	hash.Write([]byte(v))
	return hex.EncodeToString(hash.Sum(nil))
}

func EncryptWithHmacSha1(k string, v string) string {
	s := hmac.New(sha1.New, []byte(k))
	s.Write([]byte(v))
	return hex.EncodeToString(s.Sum(nil))
}

func Hmac1ToBase64(key string, str string, IsUrl bool) string {
	s := hmac.New(sha1.New, []byte(key))
	s.Write([]byte(str))
	return EncodingBase64(s.Sum(nil), IsUrl)
}

func EncodingBytesToHex(src []byte) string {
	return hex.EncodeToString(src)
}
