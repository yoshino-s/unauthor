package scanner

import (
	"context"
	"encoding/json"
	"os"
	"sync"
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
	p := pool.New().WithMaxGoroutines(s.config.Concurrent)

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

	outputLock := sync.Mutex{}

	for target := range targets {
		func(target string) {
			p.Go(func() {
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

				outputLock.Lock()
				defer outputLock.Unlock()

				if err := json.NewEncoder(os.Stdout).Encode(res); err != nil {
					panic(err)
				}
			})
		}(target)
	}

	p.Wait()
}
