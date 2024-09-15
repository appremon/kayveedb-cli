package cmd

import (
	"fmt"
	"log"

	"github.com/appremon/kayveedb-cli/client"
	"github.com/spf13/cobra"
)

func ReadCmd(c client.ClientInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "read",
		Short: "Read a value by key from KayveeDB",
		Run: func(cmd *cobra.Command, args []string) {
			readCmd := fmt.Sprintf("read %s", key)

			response, err := c.SendCommand(readCmd, hostname, port)
			if err != nil {
				log.Fatalf("Error reading key: %v", err)
			}

			fmt.Println(response)
		},
	}
}

func init() {
	RootCmd.AddCommand(ReadCmd(client.NewClient()))
	ReadCmd(client.NewClient()).Flags().StringVarP(&key, "key", "k", "", "Key to read")
	ReadCmd(client.NewClient()).MarkFlagRequired("key")
}
