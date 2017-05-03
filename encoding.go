package public

import (
	"encoding/base64"
	"net/url"
)

func ToBase64(b []byte, needUrl bool) string {
	if needUrl {
		return url.QueryEscape(base64.StdEncoding.EncodeToString(b))
	}
	return base64.StdEncoding.EncodeToString(b)
}
