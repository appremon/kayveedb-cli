package cmd

import (
	"fmt"
	"log"

	"github.com/rickcollette/kayveedb-cli/internal/client"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the KayveeDB server",
	Run: func(cmd *cobra.Command, args []string) {
		// Prompt for password if not set in environment
		password := os.Getenv("KVDB_PASS")
		if password == "" {
			fmt.Printf("%s's password: ", username)
			bytePassword, err := term.ReadPassword(0)
			if err != nil {
				log.Fatal("Failed to read password")
			}
			password = string(bytePassword)
			fmt.Println() // New line after password input
		}

		// Authenticate with the server
		err := client.Authenticate(username, password, database, hostname, port)
		if err != nil {
			log.Fatalf("Failed to login: %v", err)
		}

		fmt.Printf("Successfully logged in to %s database at %s\n", database, hostname)
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
}
