package cmd

import (
	"testing"

	"github.com/rickcollette/kayveedb-cli/mocks"
	"github.com/stretchr/testify/assert"
)

func TestReadCommand(t *testing.T) {
	mockClient := new(mocks.MockClient)

	// Set up mock expectations
	mockClient.On("SendCommand", "read key", "localhost", "3466").Return("Value for key", nil)

	// Create the read command with the mock client
	readCmd := ReadCmd(mockClient)

	// Set the flags for the read command
	readCmd.Flags().Set("key", "key")

	// Execute the command
	err := readCmd.RunE(readCmd, []string{})
	assert.Nil(t, err)

	// Check if the mock client was called
	mockClient.AssertExpectations(t)
}
