package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEnvVariableParsing tests if the CLI correctly reads environment variables
func TestEnvVariableParsing(t *testing.T) {
	// Set environment variables
	os.Setenv("USERNAME", "testuser")
	os.Setenv("PASSWORD", "testpass")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3466")

	assert.Equal(t, "testuser", getEnvOrDefault("USERNAME", ""), "USERNAME should be correctly read")
	assert.Equal(t, "testpass", getEnvOrDefault("PASSWORD", ""), "PASSWORD should be correctly read")
	assert.Equal(t, "testdb", getEnvOrDefault("DB_NAME", ""), "DB_NAME should be correctly read")
	assert.Equal(t, "localhost", getEnvOrDefault("DB_HOST", ""), "DB_HOST should be correctly read")
	assert.Equal(t, "3466", getEnvOrDefault("DB_PORT", ""), "DB_PORT should be correctly read")

	// Cleanup
	os.Unsetenv("USERNAME")
	os.Unsetenv("PASSWORD")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
}

// TestDSNParsing tests if the DSN is correctly parsed
func TestDSNParsing(t *testing.T) {
	dsn := "testuser:testpass@localhost:3466/testdb"
	parseDSN(dsn)

	assert.Equal(t, "testuser", username, "Username should be parsed from DSN")
	assert.Equal(t, "testpass", password, "Password should be parsed from DSN")
	assert.Equal(t, "localhost", hostname, "Hostname should be parsed from DSN")
	assert.Equal(t, "3466", port, "Port should be parsed from DSN")
	assert.Equal(t, "testdb", database, "Database name should be parsed from DSN")
}
