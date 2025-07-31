package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	addr := ":2801"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	bufReader := bufio.NewReader(conn)
	peek, err := bufReader.Peek(1)
	if err != nil {
		log.Printf("[SOCKS5] Failed to read from connection: %v", err)
		conn.Close()
		return
	}

	switch peek[0] {
	case 0x05:
		HandleSocks5(bufReader, conn)
	default:
		handleRestAPI(bufReader, conn)
	}
}
