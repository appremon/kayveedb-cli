package cmd

import (
	"fmt"
	"log"

	"github.com/appremon/kayveedb-cli/client"
	"github.com/spf13/cobra"
)

var key string

func DeleteCmd(c client.ClientInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a key-value pair from KayveeDB",
		Run: func(cmd *cobra.Command, args []string) {
			// Form the delete command for the server
			deleteCmd := fmt.Sprintf("delete %s", key)

			// Send the command to the server
			response, err := c.SendCommand(deleteCmd, hostname, port)
			if err != nil {
				log.Fatalf("Error deleting key: %v", err)
			}

			fmt.Println(response)
		},
	}
}

func init() {
	RootCmd.AddCommand(DeleteCmd(client.NewClient())) // Use the real client by default
}
