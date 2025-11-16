package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	HOMEPATH = "assets/html/home.html"
	PINGPATH = "assets/html/pong.html"
)

func parseRequest(raw []string) *HttpRequest {
	req := &HttpRequest{Headers: make(map[string]string)}
	if len(raw) == 0 {
		log.Printf("0 len raws")
		return nil // this seems wrong?
	}
	fline := strings.Split(raw[0], " ")
	if len(fline) != 3 {
		log.Printf("Expected first line of request to be of len 3")
		return nil
	}

	req.Method = fline[0]
	req.Path = fline[1]
	req.Proto = fline[2]

	for _, line := range raw[1:] {
		kv := strings.SplitN(line, ":", 2)
		if len(kv) != 2 {
			continue // some malformed header here
		}
		key := strings.TrimSpace(kv[0])
		val := strings.TrimSpace(kv[1])
		req.Headers[key] = val
	}

	return req
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var requestLines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("read err: ", err)
			return
		}

		trimmed := strings.TrimRight(line, "\r\n")
		if trimmed == "" {
			break
		}
		requestLines = append(requestLines, trimmed)
	}
	log.Printf("Msg receieved: %+v\n", requestLines)

	parsed := parseRequest(requestLines)
	if parsed == nil {
		log.Printf("parsing raw failed")
		return
	}
	fmt.Printf("Got request: %s %s\n", parsed.Method, parsed.Path)
	if parsed.Path == "/" {
		sendStaticFile(HOMEPATH, conn)
	} else if parsed.Path == "/ping" {
		sendStaticFile(PINGPATH, conn)
	} else if strings.HasPrefix(parsed.Path, "/assets/") {
		filePath := "." + parsed.Path
		sendStaticFile(filePath, conn)
	} else {
		sendError(conn, NOT_FOUND_HTML, NOT_FOUND_CODE)
	}
}

func main() {
	protocol := "tcp"
	port := ":8000"
	useTLS := os.Getenv("USE_TLS") == "true"
	certFile := os.Getenv("TLS_CERT_FILE")
	keyFile := os.Getenv("TLS_KEY_FILE")

	var listener net.Listener
	var err error

	if useTLS {
		if _, err := os.Stat(certFile); os.IsNotExist(err) {
			log.Printf("Certificate file %s not found, falling back to HTTP", certFile)
			useTLS = false
		}
		if _, err := os.Stat(keyFile); os.IsNotExist(err) {
			log.Printf("Key file %s not found, falling back to HTTP", keyFile)
			useTLS = false
		}

		if useTLS {
			cert, err := tls.LoadX509KeyPair(certFile, keyFile)
			if err != nil {
				log.Printf("Failed to load TLS key pair: %v", err)
				useTLS = false
			} else {
				tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
				listener, err = tls.Listen(protocol, port, tlsConfig)
				if err != nil {
					log.Printf("Failed to start TLS listener: %v", err)
					useTLS = false
				} else {
					log.Printf("Listening on HTTPS port %s", port)
				}
			}
		}
	}

	if !useTLS {
		listener, err = net.Listen(protocol, port)
		if err != nil {
			log.Fatalf("Failed to start HTTP listener: %v", err)
		}
		log.Printf("Listening on HTTP port %s", port)
	}

	log.Printf("Listening on port %s\n", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleConnection(conn)
	}

}
