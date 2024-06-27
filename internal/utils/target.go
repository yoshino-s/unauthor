package utils

import (
	"net/url"
	"strconv"
)

func ExtractAddr(target string, defaultPort int) (addr string, err error) {
	url, err := url.Parse(target)
	if err != nil {
		return
	}
	if url.Port() == "" {
		url.Host = url.Host + ":" + strconv.Itoa(defaultPort)
	}
	addr = url.Host
	return
}
