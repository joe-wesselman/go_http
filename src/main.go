package main

import (
	"bufio"
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

	listener, err := net.Listen(protocol, port)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("Listening on port %s\n", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleConnection(conn)
	}

}
