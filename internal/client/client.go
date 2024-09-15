package client

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Authenticate sends the login request to the server
func Authenticate(username, password, database, hostname, port string) error {
	conn, err := net.Dial("tcp", hostname+":"+port)
	if err != nil {
		return fmt.Errorf("failed to connect to %s:%s: %v", hostname, port, err)
	}
	defer conn.Close()

	// Send authentication command
	authCmd := fmt.Sprintf("auth %s %s %s", username, password, database)
	fmt.Fprintf(conn, authCmd+"\n")

	// Read server response
	serverReader := bufio.NewReader(conn)
	response, err := serverReader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read server response: %v", err)
	}

	// Check for authentication failure
	if strings.HasPrefix(response, "Error") {
		return fmt.Errorf("authentication failed: %s", strings.TrimSpace(response))
	}

	return nil
}

// SendCommand sends a command to the server and returns the response
func SendCommand(cmd, hostname, port string) (string, error) {
	conn, err := net.Dial("tcp", hostname+":"+port)
	if err != nil {
		return "", fmt.Errorf("failed to connect to %s:%s: %v", hostname, port, err)
	}
	defer conn.Close()

	// Send the command to the server
	fmt.Fprintf(conn, cmd+"\n")

	// Read server response
	serverReader := bufio.NewReader(conn)
	response, err := serverReader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read server response: %v", err)
	}

	return strings.TrimSpace(response), nil
}
