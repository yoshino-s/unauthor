package dubbo

import (
	"bufio"
	"context"
	"net"
	"strings"
	"time"

	"github.com/yoshino-s/unauthor/internal/scanner"
	"github.com/yoshino-s/unauthor/internal/utils"
)

var _ scanner.ScanFunc = Dubbo

const payload = "ls\n\n"

func Dubbo(ctx context.Context, target string) (res scanner.ScanFuncResult, err error) {
	res.Success = false

	var addr string

	addr, err = utils.ExtractAddr(target, 20880)

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
	var result string

	result, err = reader.ReadString('>')
	if err != nil {
		return
	}

	res.Result = result

	if err = conn.Close(); err != nil {
		return
	}

	if strings.Contains(result, "dubbo>") {
		res.Success = true
	}

	return
}
