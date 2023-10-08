package redis

import (
	"bufio"
	"context"
	"net"
	"strconv"
	"strings"

	"github.com/yoshino-s/unauthor/internal/scanner"
	"github.com/yoshino-s/unauthor/internal/utils"
)

var _ scanner.ScanFunc = Redis

// s = socket.socket()
// socket.setdefaulttimeout(3)  # 设置超时时间
// try:
// 	s.connect((host, port))
// 	s.send(payload)  # 发送info命令
// 	response = s.recv(1024).decode()
// 	s.close()
// 	if response and 'redis_version' in data:
// 		return True,'%s:%s'%(host,port)
// except (socket.error, socket.timeout):
// 	pass

const payload = "*1\r\n$4\r\ninfo\r\n"

func Redis(ctx context.Context, target string) (res scanner.ScanFuncResult, err error) {
	res.Success = false
	res.Target = target

	var addr string

	addr, err = utils.ExtractAddr(target)

	if err != nil {
		res.Error = err
		return
	}

	var conn net.Conn

	conn, err = net.Dial("tcp", addr)

	if err != nil {
		return
	}

	defer conn.Close()

	d, ok := ctx.Deadline()

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
