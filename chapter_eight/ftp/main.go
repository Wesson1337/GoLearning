package main

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("220 Welcome to the Go FTP Server\r\n"))

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Error reading from connection: %s\n", err)
			return
		}

		// Convert bytes to string and remove trailing newline
		cmd := strings.TrimSpace(string(buf[:n]))

		// Handle different FTP commands
		switch {
		case strings.HasPrefix(cmd, "cd"):
			dir, err := exec.Command("cd", cmd[2:]).Output()
			if err != nil {
				conn.Write([]byte("550 Directory command failed\r\n"))
				continue
			}
			conn.Write([]byte(fmt.Sprintf("250 Directory changed to %s\r\n", dir)))
		case strings.HasPrefix(cmd, "pwd"):
			currDir, err := exec.Command("pwd").Output()
			if err != nil {
				conn.Write([]byte("550 Directory command failed\r\n"))
				continue
			}
			conn.Write([]byte(fmt.Sprintf("257 %s\r", currDir)))
		case strings.HasPrefix(cmd, "close"):
			conn.Write([]byte("221 Goodbye!\r\n"))
			return
		default:
			conn.Write([]byte("502 Command not implemented\r\n"))
		}
	}
}

func main() {
	fmt.Println("FTP Server starting on port 2121...")

	listener, err := net.Listen("tcp", ":2121")
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
			continue
		}
		go handleConnection(conn)
	}
}
