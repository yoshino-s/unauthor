package utils

import (
	"net/url"
	"strconv"
	"strings"
)

func ExtractAddr(target string, defaultPort int) (addr string, err error) {
	var u *url.URL
	if strings.Contains(target, "://") {
		u, err = url.Parse(target)
	} else {
		u, err = url.Parse("tcp://" + target)
	}
	if err != nil {
		return
	}
	if u.Port() == "" {
		u.Host = u.Host + ":" + strconv.Itoa(defaultPort)
	}
	addr = u.Host
	return
}
