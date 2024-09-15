package cmd

import (
	"testing"

	"github.com/appremon/kayveedb-cli/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLoginCommand(t *testing.T) {
	mockClient := new(mocks.MockClient)

	// Set up mock expectations
	mockClient.On("Authenticate", "testuser", "testpass", "testdb", "localhost", "3466").Return(nil)

	// Create the login command with the mock client
	loginCmd := LoginCmd(mockClient)

	// Set environment variables or parameters
	username = "testuser"
	password = "testpass"
	database = "testdb"

	// Execute the command
	err := loginCmd.RunE(loginCmd, []string{})
	assert.Nil(t, err)

	// Check if the mock client was called
	mockClient.AssertExpectations(t)
}
