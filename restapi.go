package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
)

func handleRestAPI(r *bufio.Reader, conn net.Conn) {
	defer conn.Close()
	log.Println("[RESTAPI] Handling REST API request")

	req, err := http.ReadRequest(r)
	if err != nil {
		log.Printf("[RESTAPI] Failed to read HTTP request: %v", err)
		return
	}

	req.Close = true

	var response string
	var statusLine string

	switch req.Method {
	case http.MethodGet:
		switch req.URL.Path {
		case "/proxy":
			response = "GET request received"
			statusLine = "HTTP/1.1 200 OK\r\n"
		default:
			statusLine = "HTTP/1.1 404 Not Found\r\n"
			response = "Not Found"
		}

	case http.MethodPost:
		response = "POST request received"
		statusLine = "HTTP/1.1 200 OK\r\n"
	default:
		response = "Method not allowed"
		statusLine = "HTTP/1.1 405 Method Not Allowed\r\n"
	}

	headers := "Content-Type: text/plain\r\nContent-Length: %d\r\nConnection: close\r\n\r\n"
	headers = fmt.Sprintf(headers, len(response))
	conn.Write([]byte(statusLine + headers + response))
}
