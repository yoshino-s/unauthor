package scanner

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/sourcegraph/conc/iter"
	"github.com/sourcegraph/conc/pool"
)

type ScanFunc func(context.Context, string) (ScanFuncResult, error)

type Scanner struct {
	config   ScannerConfig
	scanFunc ScanFunc
}

func NewScanner(config ScannerConfig, scanFunc ScanFunc) *Scanner {
	return &Scanner{
		config:   config,
		scanFunc: scanFunc,
	}
}

func (s *Scanner) Run() {
	p := pool.NewWithResults[ScanFuncResult]().WithMaxGoroutines(s.config.Concurrent)

	targets := make(chan string)

	go func() {
		for _, target := range s.config.Targets {
			targets <- target
		}

		iter.ForEach(s.config.TargetsFile, func(t *string) {
			f, err := os.Open(*t)
			if err != nil {
				panic(err)
			}

			defer f.Close()

			buf := make([]byte, 1024)
			for {
				n, err := f.Read(buf)
				if err != nil {
					break
				}

				targets <- string(buf[:n])
			}
		})

		close(targets)
	}()

	for target := range targets {
		p.Go(func() ScanFuncResult {
			t := time.Now()
			ctx, cancel := context.WithTimeout(
				context.Background(),
				s.config.Timeout,
			)
			defer cancel()
			res, err := s.scanFunc(ctx, target)
			res.Target = target
			res.Time = time.Since(t)
			if err != nil {
				res.Error = err
			}
			return res
		})
	}

	res := p.Wait()

	if res == nil {
		res = []ScanFuncResult{}
	}

	r, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(r)
}
