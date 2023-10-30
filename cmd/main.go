package main

import (
	"fmt"
	"maxxfrazer/golang-proxy-server/proxy"
	"net"
	"net/http"
	"os"
)

func main() {
	serverPort, serverPortExists := os.LookupEnv("SERVER_PORT")
	if !serverPortExists || len(serverPort) == 0 {
		// Check $PORT, this is used by Railway.
		port, portExists := os.LookupEnv("PORT")
		if portExists && len(port) > 0 {
			serverPort = port
		} else {
			serverPort = "8080"
		}
	}

	// Start a simple HTTP server for health checks
	go setupHealthCheckServer()

	listener, err := net.Listen("tcp", "localhost:"+serverPort)
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Proxy server is listening on :" + serverPort)

	for {
		// Accept incoming connections
		clientConn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle incoming connection with a goroutine
		go proxy.HandleClient(clientConn)
	}
}

func setupHealthCheckServer() {
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		// Check the health of your application here
		// You can add more advanced health checks as needed

		// For a simple health check, just respond with 200 OK
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	// Listen on a separate port for health checks (e.g., :8081)
	healthCheckPort := "8081"
	fmt.Println("Health check server is listening on :" + healthCheckPort)
	err := http.ListenAndServe(":"+healthCheckPort, nil)
	if err != nil {
		fmt.Println("Error starting health check server:", err)
	}
}
