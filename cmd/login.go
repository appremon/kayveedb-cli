package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/rickcollette/kayveedb-cli/client"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func LoginCmd(c client.ClientInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Login to the KayveeDB server",
		Run: func(cmd *cobra.Command, args []string) {
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

			err := c.Authenticate(username, password, database, hostname, port)
			if err != nil {
				log.Fatalf("Failed to login: %v", err)
			}

			fmt.Printf("Successfully logged in to %s database at %s\n", database, hostname)
		},
	}
}

func init() {
	RootCmd.AddCommand(LoginCmd(client.NewClient()))
}
