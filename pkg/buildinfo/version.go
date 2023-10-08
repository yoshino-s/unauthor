package buildinfo

import (
	"fmt"
	"os"
)

var (
	Version   string = "dev"
	GitCommit string = "unknown"
	BuildTime string = "unknown"
)

func PrintVersion() {
	fmt.Printf("version: %s (%s) | %s\n", Version, GitCommit, BuildTime)
	os.Exit(0)
}
