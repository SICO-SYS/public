package public

import (
	"encoding/base64"
	"net/url"
	"strconv"
)

func EncodingBase64(b []byte, IsUrl bool) string {
	if IsUrl {
		return url.QueryEscape(base64.StdEncoding.EncodeToString(b))
	}
	return base64.StdEncoding.EncodeToString(b)
}

func StringToInt(num string) int {
	result, _ := strconv.Atoi(num)
	return result
}

func StringToFloat64(num string) float64 {
	result, _ := strconv.ParseFloat(num, 64)
	return result
}

func IntToString(num int) string {
	return strconv.Itoa(num)
}

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}
