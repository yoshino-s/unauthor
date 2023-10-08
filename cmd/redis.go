package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yoshino-s/unauthor/internal/redis"
	"github.com/yoshino-s/unauthor/internal/scanner"
)

var (
	redisCmd = &cobra.Command{
		Use:  "redis",
		Long: "Scan redis unauthorized access",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := scanner.NewScanner(config, redis.Redis)
			scanner.Run()
		},
	}
)

func init() {
	rootCmd.AddCommand(redisCmd)
}
