package app

import (
	"fmt"
	"github.com/Fantom-foundation/go-opera/version"
	"os"
	"runtime"

	"github.com/Fantom-foundation/go-opera/config"
	"gopkg.in/urfave/cli.v1"

	"github.com/Fantom-foundation/go-opera/gossip"
)

var (
	versionCommand = cli.Command{
		Action:    versionAction,
		Name:      "version",
		Usage:     "Print version numbers",
		ArgsUsage: " ",
		Category:  "MISCELLANEOUS COMMANDS",
		Description: `
The output of this command is supposed to be machine-readable.
`,
	}
)

func versionAction(ctx *cli.Context) error {
	fmt.Println(config.ClientIdentifier)
	fmt.Println("Version:", version.VersionWithMeta)
	if config.GitCommit != "" {
		fmt.Println("Git Commit:", config.GitCommit)
	}
	if config.GitDate != "" {
		fmt.Println("Git Commit Date:", config.GitDate)
	}
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Protocol Versions:", []uint{gossip.ProtocolVersion})
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Printf("GOPATH=%s\n", os.Getenv("GOPATH"))
	fmt.Printf("GOROOT=%s\n", runtime.GOROOT())
	return nil
}