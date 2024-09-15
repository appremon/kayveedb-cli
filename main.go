package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"golang.org/x/term"
)

// Command-line flags for username, database, host, and port
var (
	username string
	database string
	host     string
	port     string
	password string
)

// Function to handle user input and send commands to the server
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Set up a REPL
	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	// Continuously read from server and display prompt
	for {
		fmt.Print("kayveedb@", database, "> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}

		// Send the input to the server
		fmt.Fprintf(conn, input+"\n")

		// Read server's response
		response, _ := serverReader.ReadString('\n')
		fmt.Println(strings.TrimSpace(response))
	}
}

func main() {
	// Parse flags for CLI usage
	flag.StringVar(&username, "u", "", "Username")
	flag.StringVar(&database, "d", "", "Database name")
	flag.StringVar(&host, "h", "localhost", "Hostname or IP address of the server")
	flag.StringVar(&port, "P", "3466", "Port number")
	flag.Parse()

	// Check if password is set in environment variable
	password = os.Getenv("KVDB_PASS")

	// If no password is set, prompt for it interactively
	if password == "" {
		fmt.Printf("%s's password: ", username)
		bytePassword, err := term.ReadPassword(0)
		if err != nil {
			log.Fatal("Failed to read password")
		}
		password = string(bytePassword)
		fmt.Println() // Print a newline after password input
	}

	// Connect to the server
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("Failed to connect to %s:%s: %v", host, port, err)
	}
	defer conn.Close()

	// Send the authentication command
	authCmd := fmt.Sprintf("auth %s %s %s", username, password, database)
	fmt.Fprintf(conn, authCmd+"\n")

	// Wait for server response
	serverReader := bufio.NewReader(conn)
	serverResponse, _ := serverReader.ReadString('\n')
	if strings.HasPrefix(serverResponse, "Error") {
		log.Fatalf("Authentication failed: %s", strings.TrimSpace(serverResponse))
	}

	// Server response is successful, proceed to handle REPL commands
	fmt.Printf("Connected to %s database on %s\n", database, host)
	handleConnection(conn)
}
