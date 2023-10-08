package utils

import (
	"net/url"
	"strings"
)

func ExtractAddr(target string) (addr string, err error) {
	if strings.Contains(target, "//") {
		var url *url.URL
		url, err = url.Parse(target)
		if err != nil {
			return
		}
		addr = url.Host
		return
	} else {
		addr = target
		return
	}
}
