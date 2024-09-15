package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// TestVersionCmd tests the version command output
func TestVersionCmd(t *testing.T) {
	// Set up a buffer to capture command output
	var buf bytes.Buffer
	VersionCmd.SetOut(&buf)

	// Execute the version command
	err := VersionCmd.RunE(&cobra.Command{}, []string{})
	assert.Nil(t, err)

	// Check the output
	expected := "kvdbcli version: v1.0.2\n"
	assert.Equal(t, expected, buf.String())
}
