package types

import (
	"context"
	"time"
)

type ScanFuncResult struct {
	Target  string        `json:"target"`
	Success bool          `json:"success"`
	Result  string        `json:"result"`
	Error   string        `json:"error"`
	Time    time.Duration `json:"time"`
}

type ScanFunc func(context.Context, string) (ScanFuncResult, error)
