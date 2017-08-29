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

func String2Int(num string) int {
	result, _ := strconv.Atoi(num)
	return result
}

func String2Float(num string) float64 {
	result, _ := strconv.ParseFloat(num, 64)
	return result
}

func Int2String(num int) string {
	return strconv.Itoa(num)
}

func Int642String(num int64) string {
	return strconv.FormatInt(num, 10)
}

func Int2Float(num int) float64 {
	return float64(num)
}
