package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rickcollette/kayveedb-cli/client"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	username string
	password string
	database string
	hostname string
	port     string
)

// RootCmd is the base command for the CLI.
var RootCmd = &cobra.Command{
	Use:   "kayveedb-cli",
	Short: "CLI to interact with KayveeDB server",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// First check for DB_DSN environment variable
		dsn := os.Getenv("DB_DSN")
		if dsn != "" {
			parseDSN(dsn)
		} else {
			// Fallback to individual environment variables
			username = getEnvOrDefault("USERNAME", username)
			password = getEnvOrDefault("PASSWORD", password)
			database = getEnvOrDefault("DB_NAME", database)
			hostname = getEnvOrDefault("DB_HOST", hostname)
			port = getEnvOrDefault("DB_PORT", port)

			// If any essential info is still missing, prompt for it
			if username == "" {
				fmt.Print("username: ")
				fmt.Scanln(&username)
			}
			if database == "" {
				fmt.Print("database: ")
				fmt.Scanln(&database)
			}
			if password == "" {
				fmt.Printf("%s's password: ", username)
				bytePassword, err := term.ReadPassword(0)
				if err != nil {
					log.Fatal("Failed to read password")
				}
				password = string(bytePassword)
				fmt.Println() // New line after password input
			}
		}

		// Create a new client and authenticate
		realClient := client.NewClient()

		err := realClient.Authenticate(username, password, database, hostname, port)
		if err != nil {
			log.Fatalf("Failed to login: %v", err)
		}
	},
}

// Execute runs the root command and its subcommands.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username for the KayveeDB server")
	RootCmd.PersistentFlags().StringVarP(&database, "database", "d", "", "Database name")
	RootCmd.PersistentFlags().StringVarP(&hostname, "hostname", "h", "localhost", "Hostname of the KayveeDB server")
	RootCmd.PersistentFlags().StringVarP(&port, "port", "P", "3466", "Port of the KayveeDB server")
}

// Helper function to get environment variable or default to the provided value
func getEnvOrDefault(envKey, defaultValue string) string {
	if value := os.Getenv(envKey); value != "" {
		return value
	}
	return defaultValue
}

// parseDSN parses a DSN string and sets the connection parameters
func parseDSN(dsn string) {
	// Example DSN format: "username:password@hostname:port/database"
	parts := strings.Split(dsn, "@")
	if len(parts) != 2 {
		log.Fatalf("Invalid DB_DSN format")
	}

	// Parse the credentials part (username:password)
	credentials := strings.Split(parts[0], ":")
	if len(credentials) != 2 {
		log.Fatalf("Invalid DB_DSN credentials format")
	}
	username = credentials[0]
	password = credentials[1]

	// Parse the host part (hostname:port/database)
	hostPart := parts[1]
	hostParts := strings.Split(hostPart, "/")
	if len(hostParts) != 2 {
		log.Fatalf("Invalid DB_DSN hostname/database format")
	}
	database = hostParts[1]

	hostAndPort := strings.Split(hostParts[0], ":")
	if len(hostAndPort) != 2 {
		log.Fatalf("Invalid DB_DSN hostname:port format")
	}
	hostname = hostAndPort[0]
	port = hostAndPort[1]
}
