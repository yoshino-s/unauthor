package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/yoshino-s/go-framework/application"
	"github.com/yoshino-s/go-framework/cmd"
	"github.com/yoshino-s/go-framework/common"
	"github.com/yoshino-s/go-framework/configuration"
	"github.com/yoshino-s/unauthor/scanner"
)

var name = "unauthor"
var app = application.NewMainApplication()

var (
	scannerApp = scanner.New()
	rootCmd    = &cobra.Command{
		Use: name,
		Run: func(cmd *cobra.Command, args []string) {
			app.Append(scannerApp)
			app.Go(context.Background())
		},
	}
)

func init() {
	common.AppName = name

	cobra.OnInitialize(func() {
		configuration.Setup(name)
	})

	configuration.GenerateConfiguration.Register(rootCmd.PersistentFlags())
	app.Configuration().Register(rootCmd.PersistentFlags())
	scannerApp.Configuration().Register(rootCmd.PersistentFlags())

	rootCmd.AddCommand(cmd.VersionCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
