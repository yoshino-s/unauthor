package memcached

import (
	"bufio"
	"context"
	"net"
	"strings"
	"time"

	"github.com/yoshino-s/unauthor/scanner/types"
	"github.com/yoshino-s/unauthor/utils"
)

var _ types.ScanFunc = Memcached

const payload = "stats\r\n"

func Memcached(ctx context.Context, target string) (res types.ScanFuncResult, err error) {
	res.Success = false

	var addr string

	addr, err = utils.ExtractAddr(target, 11211)

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

	reader := bufio.NewScanner(conn)
	var result string

	for reader.Scan() {
		text := reader.Text() + "\n"
		result += text
		if strings.Contains(text, "END") {
			break
		}
	}

	res.Result = result

	if err = reader.Err(); err != nil {
		return
	}

	if err = conn.Close(); err != nil {
		return
	}

	if strings.Contains(result, "version") {
		res.Success = true
	}

	return
}
