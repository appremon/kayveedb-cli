package cmd

import (
	"testing"

	"github.com/appremon/kayveedb-cli/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCommand(t *testing.T) {
	mockClient := new(mocks.MockClient)

	// Set up mock expectations
	mockClient.On("SendCommand", "update key=newvalue", "localhost", "3466").Return("Update successful", nil)

	// Create the update command with the mock client
	updateCmd := UpdateCmd(mockClient)

	// Set the flags for the update command
	updateCmd.Flags().Set("key", "key")
	updateCmd.Flags().Set("value", "newvalue")

	// Execute the command
	err := updateCmd.RunE(updateCmd, []string{})
	assert.Nil(t, err)

	// Check if the mock client was called
	mockClient.AssertExpectations(t)
}
