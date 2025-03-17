package ftp

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/yoshino-s/unauthor/scanner/types"
	"github.com/yoshino-s/unauthor/utils"
)

var _ types.ScanFunc = Ftp

func Ftp(ctx context.Context, target string) (res types.ScanFuncResult, err error) {
	res.Success = false

	addr, err := utils.ExtractAddr(target, 21)

	if err != nil {
		res.Error = err.Error()
		return
	}

	d, ok := ctx.Deadline()

	var opts []ftp.DialOption
	if ok {
		opts = append(opts, ftp.DialWithTimeout(time.Until(d)))
	}
	c, err := ftp.Dial(addr.Host, opts...)
	if err != nil {
		res.Error = err.Error()
		return
	}

	defer c.Quit()

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

	res.Exploit = fmt.Sprintf("echo -e 'open %s %s\nuser anonymous anonymous\nls\nquit' | ftp -n", addr.Hostname(), addr.Port())
	res.Result = strings.Join(files, "\n")
	res.Success = true

	return
}
