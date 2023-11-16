package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/yoshino-s/unauthor/internal/redis"
	"github.com/yoshino-s/unauthor/internal/scanner"
)

var (
	rootCmd = &cobra.Command{
		Use: "unauthor",
		Run: func(cmd *cobra.Command, args []string) {
			var s *scanner.Scanner
			switch scanType {
			case "redis":
				s = scanner.NewScanner(config, redis.Redis)
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
	rootCmd.PersistentFlags().StringVar(&scanType, "type", "", "scan type")
	rootCmd.PersistentFlags().StringArrayVarP(&config.Targets, "targets", "t", []string{}, "target files or directories")
	rootCmd.PersistentFlags().StringArrayVarP(&config.TargetsFile, "targets-file", "f", []string{}, "target files or directories")
	rootCmd.PersistentFlags().DurationVarP(&config.Timeout, "timeout", "T", time.Second*10, "timeout seconds")
	rootCmd.PersistentFlags().IntVarP(&config.Concurrent, "concurrent", "c", 20, "concurrent number")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
