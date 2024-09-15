package cmd

import (
	"testing"

	"github.com/rickcollette/kayveedb-cli/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCommand(t *testing.T) {
	// Create a new mock client
	mockClient := new(mocks.MockClient)

	// Set up mock expectations
	mockClient.On("SendCommand", "delete key", "localhost", "3466").Return("Delete successful", nil)

	// Create the delete command with the mock client
	deleteCmd := DeleteCmd(mockClient)

	// Set the flags for the delete command
	deleteCmd.Flags().Set("key", "key")

	// Execute the command
	err := deleteCmd.RunE(deleteCmd, []string{})
	assert.Nil(t, err)

	// Check if the mock client method was called
	mockClient.AssertExpectations(t)
}
