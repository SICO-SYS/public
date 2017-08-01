package public

import (
	"encoding/base64"
	"net/url"
	"strconv"
)

func ToBase64(b []byte, needUrl bool) string {
	if needUrl {
		return url.QueryEscape(base64.StdEncoding.EncodeToString(b))
	}
	return base64.StdEncoding.EncodeToString(b)
}

func Atoi(num string) int {
	result, _ := strconv.Atoi(num)
	return result
}

func Itoa(num int) string {
	return strconv.Itoa(num)
}
