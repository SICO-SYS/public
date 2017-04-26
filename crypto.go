/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package public

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func Sha256Encrypt(v string) string {
	hash := sha256.New()
	hash.Write([]byte(v))
	return hex.EncodeToString(hash.Sum(nil))
}

func HMACSha256Encrypt(k string, v string) string {
	s := hmac.New(sha256.New, []byte(k))
	s.Write([]byte(v))
	return hex.EncodeToString(s.Sum(nil))
}