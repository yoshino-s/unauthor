package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/yoshino-s/unauthor/internal/dubbo"
	"github.com/yoshino-s/unauthor/internal/memcached"
	"github.com/yoshino-s/unauthor/internal/redis"
	"github.com/yoshino-s/unauthor/internal/scanner"
	"github.com/yoshino-s/unauthor/internal/zookeeper"
)

var (
	rootCmd = &cobra.Command{
		Use: "unauthor",
		Run: func(cmd *cobra.Command, args []string) {
			var s *scanner.Scanner
			switch scanType {
			case "redis":
				s = scanner.NewScanner(config, redis.Redis)
			case "zookeeper":
				s = scanner.NewScanner(config, zookeeper.Zookeeper)
			case "memcached":
				s = scanner.NewScanner(config, memcached.Memcached)
			case "dubbo":
				s = scanner.NewScanner(config, dubbo.Dubbo)
			case "jdwp":
				s = scanner.NewScanner(config, dubbo.Dubbo)
			default:
				cobra.CheckErr("unknown scan type")
			}
			s.Run()
		},
	}
	config   scanner.ScannerConfig
	scanType string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&scanType, "type", "", "scan type, one of redis, zookeeper, memcached, dubbo, jdwp")
	rootCmd.PersistentFlags().StringSliceVarP(&config.Targets, "targets", "t", []string{}, "target files or directories")
	rootCmd.PersistentFlags().StringSliceVarP(&config.TargetsFile, "targets-file", "f", []string{}, "target files or directories")
	rootCmd.PersistentFlags().DurationVarP(&config.Timeout, "timeout", "T", time.Second*10, "timeout seconds")
	rootCmd.PersistentFlags().IntVarP(&config.Concurrent, "concurrent", "c", 20, "concurrent number")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
