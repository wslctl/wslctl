package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Microsoft/go-winio"
)

func main() {
	// Define the named pipe you're connecting to
	pipeName := `\\.\pipe\wslctl`

	// Dial the named pipe using winio
	conn, err := winio.DialPipe(pipeName, nil)
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Send a command
	_, err = fmt.Fprintf(conn, "%s\n", os.Args[1])
	if err != nil {
		fmt.Printf("Failed to send message: %v\n", err)
		os.Exit(1)
	}

	// Read the response from the pipe
	reader := bufio.NewReader(conn)
	response, _ := reader.ReadString('\n')
	fmt.Println("Response:", response)
}
