package scanner

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/yoshino-s/go-framework/common"
	"github.com/yoshino-s/go-framework/configuration"
	"github.com/yoshino-s/unauthor/scanner/types"
)

var _ configuration.Configuration = (*ScannerConfig)(nil)

type ScannerConfig struct {
	Targets     []string
	TargetsFile []string
	Timeout     time.Duration
	Concurrent  int
	Protocol    string
	ScanFuncs   map[string]types.ScanFunc
}

func (c *ScannerConfig) Register(set *pflag.FlagSet) {
	set.StringSliceVarP(&c.Targets, "scanner.targets", "t", nil, "targets")
	set.StringSliceVarP(&c.TargetsFile, "scanner.targets-file", "f", nil, "targets file")
	set.DurationVarP(&c.Timeout, "scanner.timeout", "T", 5*time.Second, "timeout")
	set.IntVarP(&c.Concurrent, "scanner.concurrent", "c", 100, "concurrent")
	set.StringVarP(&c.Protocol, "scanner.protocol", "p", "", fmt.Sprintf("protocol, one of %s", strings.Join(common.MapKeys(c.ScanFuncs), ", ")))

	common.MustNoError(viper.BindPFlags(set))
	configuration.Register(c)
}

func (c *ScannerConfig) Read() {
	common.MustDecodeFromMapstructure(viper.AllSettings()["scanner"], c)

}
