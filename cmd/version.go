package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yitume/shop-api/pkg/version"
)

var short bool

// startCmd represents the hello command
var versionCmd = &cobra.Command{
	Use:  "version",
	Long: `Show version information`,
	Run:  versionFn,
}

func init() {
	versionCmd.PersistentFlags().BoolVarP(&short, "short", "s", short, "Displays a short form of the version information")
	RootCmd.AddCommand(versionCmd)
}

func versionFn(cmd *cobra.Command, args []string) {
	if short {
		fmt.Println(version.Info)
	} else {
		fmt.Println(version.Info.LongForm())
	}
}
