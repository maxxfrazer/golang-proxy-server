# Simple Golang Proxy Server

A simple proxy server implemented in Golang that allows you to route incoming traffic to a specified destination server based on the provided URL and port.

## Features

- Basic TCP proxy server that forwards traffic to a destination server.
- Customizable proxy server address and port.
- Easy-to-use for routing traffic from a client to a specified destination.


## Getting Started

### One-click launch options

| Railway |
|:-:|
| [![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/template/y0AKsJ?referralCode=waRWUT) |

### Prerequisites

Before you begin, ensure you have a Go development environment set up.

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/maxxfrazer/golang-proxy-server.git
    cd golang-proxy-server
    ```

2. Build the project:

    ```sh
    go build cmd/main.go
    ```

## Usage

1. Run the proxy server:

    ```sh
    ./main
    ```
    The proxy server will start listening on localhost:8080 by default. You can customize the address and port in the main.go file.

2. Connect to the proxy server from a client application (e.g., telnet or netcat) and send a request with the target URL and port you want to access. For example:

    ```sh
    echo "GET http://example.com:80 HTTP/1.1" | nc localhost 8080
    ```

    This command sends an HTTP request to your proxy server, asking it to connect to `example.com` on port 80.

3. The proxy server will forward the request to the target server (example.com) and return the response to the client.


## Configuration

You can customize the proxy server's behavior and settings by modifying the `main.go` file. For example, you can change the proxy server's listening address and port, add authentication, or implement error handling and logging as needed.

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please fork the repository and create a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

