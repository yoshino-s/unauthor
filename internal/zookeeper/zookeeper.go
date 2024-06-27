package zookeeper

import (
	"context"
	"net"
	"strings"
	"time"

	"github.com/yoshino-s/unauthor/internal/scanner"
	"github.com/yoshino-s/unauthor/internal/utils"
)

var _ scanner.ScanFunc = Zookeeper

const payload = "envi"

func Zookeeper(ctx context.Context, target string) (res scanner.ScanFuncResult, err error) {
	res.Success = false

	var addr string

	addr, err = utils.ExtractAddr(target, 2181)

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

	var buf [1024]byte
	var result string

	for {
		n, err := conn.Read(buf[:])
		result += string(buf[:n])
		if err != nil {
			break
		}
		if n == 0 {
			break
		}
	}

	conn.Close()

	res.Result = result

	if strings.Contains(result, "Environment") {
		res.Success = true

	}
	return
}
