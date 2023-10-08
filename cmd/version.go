package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yoshino-s/unauthor/pkg/buildinfo"
)

var (
	versionCmd = &cobra.Command{
		Use:  "version",
		Long: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd.Name())
			buildinfo.PrintVersion()
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
