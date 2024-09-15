package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version string = "v1.0.3"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version of kvdbcli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kvdbcli version: %s\n", Version)
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}
