package proxy

import (
	"fmt"
	"io"
	"net"
	"net/url"
)

func HandleClient(clientConn net.Conn) {
	// Read the URL and port from the client
	buffer := make([]byte, 1024)
	_, err := clientConn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading client request:", err)
		clientConn.Close()
		return
	}

	// Parse the URL from the client request
	clientURL := string(buffer)
	targetURL, err := url.Parse(clientURL)
	if err != nil {
		fmt.Println("Error parsing client URL:", err)
		clientConn.Close()
		return
	}

	// Connect to the destination server
	destServer, err := net.Dial("tcp", targetURL.Host)
	if err != nil {
		fmt.Println("Error connecting to destination server:", err)
		clientConn.Close()
		return
	}
	defer destServer.Close()

	// Copy data from client to destination and vice versa
	go io.Copy(destServer, clientConn)
	go io.Copy(clientConn, destServer)
}
