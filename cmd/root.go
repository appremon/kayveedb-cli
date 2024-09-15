package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Global variables to store flags
var username, database, hostname, port string

// RootCmd defines the base command
var RootCmd = &cobra.Command{
	Use:   "kayveedb-cli",
	Short: "KayveeDB CLI tool to interact with KayveeDB server",
}

// Execute adds all child commands to the root command and sets flags appropriately
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Global flags for the CLI
	RootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username for login")
	RootCmd.PersistentFlags().StringVarP(&database, "database", "d", "", "Database name")
	RootCmd.PersistentFlags().StringVarP(&hostname, "hostname", "h", "localhost", "Hostname or IP address of the server")
	RootCmd.PersistentFlags().StringVarP(&port, "port", "P", "3466", "Port number")

	RootCmd.MarkPersistentFlagRequired("username")
	RootCmd.MarkPersistentFlagRequired("database")
}
