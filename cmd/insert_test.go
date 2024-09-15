package cmd

import (
	"testing"

	"github.com/appremon/kayveedb-cli/mocks"
	"github.com/stretchr/testify/assert"
)

func TestInsertCommand(t *testing.T) {
	mockClient := new(mocks.MockClient)

	// Set up mock expectations
	mockClient.On("SendCommand", "insert key=value", "localhost", "3466").Return("Insert successful", nil)

	// Create the insert command with the mock client
	insertCmd := InsertCmd(mockClient)

	// Set the flags for the insert command
	insertCmd.Flags().Set("key", "key")
	insertCmd.Flags().Set("value", "value")

	// Execute the command
	err := insertCmd.RunE(insertCmd, []string{})
	assert.Nil(t, err)

	// Check if the mock client was called
	mockClient.AssertExpectations(t)
}
