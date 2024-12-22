package auth

import (
	"errors"
	"net/http"
	"strings"
)

// getapikey extract the api key
// from header authorization of a http request
// header autorization format: {key : value }
// Authorization : ApiKey {api string here}

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("malformed authentication info found")
	}
	return vals[1], nil
}
