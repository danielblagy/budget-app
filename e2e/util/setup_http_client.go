package util

import (
	"net/http"
)

func SetupHttpClient() *http.Client {
	return &http.Client{Jar: NewCookieJar()}
}
