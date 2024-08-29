package scanner

import (
	"bufio"
	"context"
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/sourcegraph/conc/iter"
	"github.com/sourcegraph/conc/pool"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/configuration"
	"github.com/yoshino-s/unauthor/scanner/dubbo"
	"github.com/yoshino-s/unauthor/scanner/ftp"
	"github.com/yoshino-s/unauthor/scanner/jdwp"
	"github.com/yoshino-s/unauthor/scanner/memcached"
	"github.com/yoshino-s/unauthor/scanner/redis"
	"github.com/yoshino-s/unauthor/scanner/types"
	"github.com/yoshino-s/unauthor/scanner/zookeeper"
	"go.uber.org/zap"
)

type OutputFunc func(res types.ScanFuncResult)

func DefaultOutputFunc() OutputFunc {
	outputLock := sync.Mutex{}

	return func(res types.ScanFuncResult) {
		outputLock.Lock()
		defer outputLock.Unlock()

		if err := json.NewEncoder(os.Stdout).Encode(res); err != nil {
			panic(err)
		}
	}
}

type Scanner struct {
	*application.EmptyApplication
	config     ScannerConfig
	OutputFunc OutputFunc
}

func New() *Scanner {
	s := &Scanner{
		EmptyApplication: application.NewEmptyApplication(),
		config: ScannerConfig{
			scanFuncs: make(map[string]types.ScanFunc),
		},
	}

	s.Register("zookeeper", zookeeper.Zookeeper)
	s.Register("redis", redis.Redis)
	s.Register("memcached", memcached.Memcached)
	s.Register("ftp", ftp.Ftp)
	s.Register("dubbo", dubbo.Dubbo)
	s.Register("jdwp", jdwp.Jdwp)

	s.OutputFunc = DefaultOutputFunc()

	return s
}

func (s *Scanner) Register(name string, fun types.ScanFunc) {
	s.config.scanFuncs[name] = fun
}

func (s *Scanner) Configuration() configuration.Configuration {
	return &s.config
}

func (s *Scanner) Run(ctx context.Context) {
	p := pool.New().WithMaxGoroutines(s.config.Concurrent)

	scanFunc := s.config.scanFuncs[s.config.Protocol]
	if scanFunc == nil {
		s.Logger.Fatal("unknown protocol", zap.String("protocol", s.config.Protocol))
	}

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
			scanner := bufio.NewScanner(f)

			defer f.Close()

			for scanner.Scan() {
				targets <- scanner.Text()
			}
		})

		close(targets)
	}()

	for target := range targets {
		func(target string) {
			p.Go(func() {
				t := time.Now()
				ctx, cancel := context.WithTimeout(
					context.Background(),
					s.config.Timeout,
				)
				defer cancel()

				res, err := scanFunc(ctx, target)
				res.Target = target
				res.Time = time.Since(t)
				if err != nil {
					res.Error = err.Error()
				}

				s.OutputFunc(res)
			})
		}(target)
	}

	p.Wait()
}
