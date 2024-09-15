package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRootCmdInitialization checks if the root command is correctly initialized.
func TestRootCmdInitialization(t *testing.T) {
	assert.NotNil(t, RootCmd, "Root command should not be nil")
	assert.Equal(t, "kayveedb-cli", RootCmd.Use, "Root command should have the correct use")
}

// TestRootCmdFlags ensures the root command has the necessary flags
func TestRootCmdFlags(t *testing.T) {
	cmd := RootCmd

	flag := cmd.Flag("username")
	assert.NotNil(t, flag, "Username flag should exist")

	flag = cmd.Flag("database")
	assert.NotNil(t, flag, "Database flag should exist")

	flag = cmd.Flag("hostname")
	assert.NotNil(t, flag, "Hostname flag should exist")

	flag = cmd.Flag("port")
	assert.NotNil(t, flag, "Port flag should exist")
}
