package ftp

import (
	"context"
	"net"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/yoshino-s/unauthor/scanner/types"
	"github.com/yoshino-s/unauthor/utils"
)

var _ types.ScanFunc = Ftp

func Ftp(ctx context.Context, target string) (res types.ScanFuncResult, err error) {
	res.Success = false

	var addr string

	addr, err = utils.ExtractAddr(target, 21)

	if err != nil {
		res.Error = err.Error()
		return
	}

	var conn net.Conn
	d, ok := ctx.Deadline()

	var opts []ftp.DialOption
	if ok {
		opts = append(opts, ftp.DialWithTimeout(time.Until(d)))
	}
	c, err := ftp.Dial(addr, opts...)
	if err != nil {
		res.Error = err.Error()
		return
	}

	defer c.Quit()

	if ok {
		conn.SetDeadline(d)
	}

	err = c.Login("anonymous", "anonymous")
	if err != nil {
		res.Error = err.Error()
		return
	}

	entries, err := c.List("/")
	if err != nil {
		res.Error = err.Error()
		return
	}

	var files []string
	for _, entry := range entries {
		files = append(files, entry.Name)
	}
	res.Result = files
	res.Success = true

	return
}
