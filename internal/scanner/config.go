package scanner

import (
	"time"
)

type ScannerConfig struct {
	Targets     []string
	TargetsFile []string
	Timeout     time.Duration
	Concurrent  int
}

type ScanFuncResult struct {
	Target  string
	Success bool
	Result  interface{}
	Error   error
	Time    time.Duration `json:"time"`
}
