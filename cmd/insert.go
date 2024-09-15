package cmd

import (
	"fmt"
	"log"

	"github.com/rickcollette/kayveedb-cli/client"
	"github.com/spf13/cobra"
)

var value string

func InsertCmd(c client.ClientInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "insert",
		Short: "Insert a key-value pair into KayveeDB",
		Run: func(cmd *cobra.Command, args []string) {
			insertCmd := fmt.Sprintf("insert %s=%s", key, value)

			response, err := c.SendCommand(insertCmd, hostname, port)
			if err != nil {
				log.Fatalf("Error inserting key: %v", err)
			}

			fmt.Println(response)
		},
	}
}

func init() {
	RootCmd.AddCommand(InsertCmd(client.NewClient()))
	InsertCmd(client.NewClient()).Flags().StringVarP(&key, "key", "k", "", "Key to insert")
	InsertCmd(client.NewClient()).Flags().StringVarP(&value, "value", "v", "", "Value to insert")
	InsertCmd(client.NewClient()).MarkFlagRequired("key")
	InsertCmd(client.NewClient()).MarkFlagRequired("value")
}
