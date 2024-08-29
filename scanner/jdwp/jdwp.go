package jdwp

import (
	"context"
	"encoding/json"
	"net"
	"time"

	"github.com/yoshino-s/unauthor/scanner/types"
	"github.com/yoshino-s/unauthor/utils"
)

var _ types.ScanFunc = Jdwp

func Jdwp(ctx context.Context, target string) (res types.ScanFuncResult, err error) {
	res.Success = false

	var addr string

	addr, err = utils.ExtractAddr(target, 8000)

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

	jdwpConn := NewJdwpConn(conn)

	if err = jdwpConn.Handshake(); err != nil {
		res.Error = err.Error()
		return
	}

	var version *JdwpVersionResponse
	version, err = jdwpConn.Version()
	if err != nil {
		res.Error = err.Error()
		return
	}

	versionStr, err := json.Marshal(version)
	if err != nil {
		res.Error = err.Error()
		return
	}

	res.Success = true

	res.Result = string(versionStr)

	return
}
