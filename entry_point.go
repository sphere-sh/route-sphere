package route_sphere

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"errors"
	"log"
	"net"
	"strings"
	"sync"
)

type EntryPoint struct {
	Name    string   `yaml:"name"`
	Address string   `yaml:"address"`
	Domains []string `yaml:"domains"`
}

func (ep EntryPoint) VerifyConfiguration() error {
	if ep.Address == "" {
		return errors.New("address is required")
	}
	return nil
}

func (ep EntryPoint) Start(wg *sync.WaitGroup) {
	defer wg.Done()

	cert, err := tls.LoadX509KeyPair("localhost.pem", "localhost-key.pem")
	if err != nil {
		log.Fatalf("failed to load TLS certificate: %v", err)
	}
	cert2, err := tls.LoadX509KeyPair("sub.client.x.local.pem", "sub.client.x.local-key.pem")
	if err != nil {
		log.Fatalf("failed to load TLS certificate: %v", err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert, cert2}}

	listener, err := tls.Listen("tcp", ep.Address, config)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", ep.Address, err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %v", err)
			continue
		}
		go ep.handleConnection(conn)
	}
}

func (ep EntryPoint) handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Printf("failed to read from connection: %v", err)
		return
	}

	domain := getDomainFromRequest(buf)
	if domain == "" {
		sendResponse(conn, "400 Bad Request", "Bad Request")
		return
	}

	if !ep.servesDomain(domain) {
		sendResponse(conn, "404 Not Found", "Not Found")
		return
	}

	sendResponse(conn, "200 OK", "Hello, World!")
}

func getDomainFromRequest(buf []byte) string {
	reader := bufio.NewReader(bytes.NewReader(buf))
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return ""
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		if strings.HasPrefix(line, "Host:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "Host:"))
		}
	}
	return ""
}

func (ep EntryPoint) servesDomain(domain string) bool {
	for _, d := range ep.Domains {
		if domain == d {
			return true
		}
	}
	return false
}

func sendResponse(conn net.Conn, status, body string) {
	conn.Write([]byte("HTTP/1.1 " + status + "\r\n"))
	conn.Write([]byte("Content-Type: text/plain\r\n"))
	conn.Write([]byte("Connection: close\r\n"))
	conn.Write([]byte("\r\n"))
	conn.Write([]byte(body))
}
