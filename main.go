package main

import (
	"bufio"
	"fmt"
	"log"

	"github.com/Microsoft/go-winio"
)

func main() {
	conn, err := winio.DialPipe(`\\.\pipe\wslctl`, nil)
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "hello\n")

	resp, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(resp) // prints "hello from wslctl-server"
}
