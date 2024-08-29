package redis

import (
	"bufio"
	"context"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/yoshino-s/unauthor/scanner/types"
	"github.com/yoshino-s/unauthor/utils"
)

var _ types.ScanFunc = Redis

const payload = "*1\r\n$4\r\ninfo\r\n"

func Redis(ctx context.Context, target string) (res types.ScanFuncResult, err error) {
	res.Success = false

	var addr string

	addr, err = utils.ExtractAddr(target, 6379)

	if err != nil {
		res.Error = err.Error()
		return
	}

	var conn net.Conn
	d, ok := ctx.Deadline()

	if ok {
		conn, err = net.DialTimeout("tcp", addr, time.Until(d))
	} else {
		conn, err = net.Dial("tcp", addr)
	}

	if err != nil {
		return
	}

	defer conn.Close()

	if ok {
		conn.SetDeadline(d)
	}

	_, err = conn.Write([]byte(payload))

	if err != nil {
		return
	}

	reader := bufio.NewReader(conn)

	l, _, err := reader.ReadLine()

	if err != nil {
		return
	}

	if len(l) < 2 || l[0] != '$' {
		return
	}

	len, err := strconv.Atoi(string(l[1:]))

	if err != nil {
		return
	}

	if len <= 0 {
		return
	}

	// read response

	buf := make([]byte, len+2)

	totalRead := 0
	var n int

	for {
		n, err = reader.Read(buf[totalRead:])

		if err != nil {
			return
		}

		totalRead += n

		if totalRead >= len+2 {
			break
		}
	}

	conn.Close()

	r := string(buf[:len])

	res.Result = r

	if strings.Contains(r, "redis_version") {
		res.Success = true

	}
	return
}
