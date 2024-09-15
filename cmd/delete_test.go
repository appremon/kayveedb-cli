package cmd

import (
	"testing"

	"github.com/appremon/kayveedb-cli/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCommand(t *testing.T) {
	// Create a new mock client
	mockClient := new(mocks.MockClient)

	// Set up mock expectations
	mockClient.On("SendCommand", "delete key", "localhost", "3466").Return("Delete successful", nil)

	// Set the key variable to avoid nil pointer dereference
	key = "key"

	// Create the delete command with the mock client
	deleteCmd := DeleteCmd(mockClient)

	// Execute the command
	err := deleteCmd.RunE(deleteCmd, []string{})
	assert.Nil(t, err)

	// Check if the mock client method was called
	mockClient.AssertExpectations(t)
}
