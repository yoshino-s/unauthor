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
	Target  string        `json:"target"`
	Success bool          `json:"success"`
	Result  interface{}   `json:"result"`
	Error   error         `json:"error"`
	Time    time.Duration `json:"time"`
}
