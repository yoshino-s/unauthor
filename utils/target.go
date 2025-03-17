package utils

import (
	"net/url"
	"strconv"
	"strings"
)

func ExtractAddr(target string, defaultPort int) (*url.URL, error) {
	var u *url.URL
	var err error
	if strings.Contains(target, "://") {
		u, err = url.Parse(target)
	} else {
		u, err = url.Parse("tcp://" + target)
	}
	if err != nil {
		return nil, err
	}
	if u.Port() == "" {
		u.Host = u.Host + ":" + strconv.Itoa(defaultPort)
	}
	return u, nil
}
