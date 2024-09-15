package cmd

import (
	"fmt"
	"log"

	"github.com/appremon/kayveedb-cli/client"
	"github.com/spf13/cobra"
)

func UpdateCmd(c client.ClientInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update a value by key in KayveeDB",
		Run: func(cmd *cobra.Command, args []string) {
			updateCmd := fmt.Sprintf("update %s=%s", key, value)

			response, err := c.SendCommand(updateCmd, hostname, port)
			if err != nil {
				log.Fatalf("Error updating key: %v", err)
			}

			fmt.Println(response)
		},
	}
}

func init() {
	RootCmd.AddCommand(UpdateCmd(client.NewClient()))
	UpdateCmd(client.NewClient()).Flags().StringVarP(&key, "key", "k", "", "Key to update")
	UpdateCmd(client.NewClient()).Flags().StringVarP(&value, "value", "v", "", "New value to update")
	UpdateCmd(client.NewClient()).MarkFlagRequired("key")
	UpdateCmd(client.NewClient()).MarkFlagRequired("value")
}
