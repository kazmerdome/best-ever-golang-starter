package misc

import "net/http"

// GetHeaderString ...
func GetHeaderString(key string, headers http.Header) string {
	var value string
	for headerName, headerValue := range headers {
		if headerName == key {
			value = headerValue[0]
		}
	}
	return value
}
