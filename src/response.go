package main

import (
	"log"
	"net"
)

const (
	INTERNAL_SERVICE_ERR_HTML = "<h1>500 Internal Server Error</h1>"
	NOT_FOUND_HTML            = "<h1>404 File Not Found</h1>"
	BAD_REQUEST_HTML          = "<h1>400 Bad Request</h1>"
	NOT_FOUND_CODE            = 404
	INTERNAL_SERVICE_ERR_CODE = 500
	BAD_REQUEST_ERR_CODE      = 400
	OK_CODE                   = 200
)

func sendError(conn net.Conn, body string, status_code int) {
	headers := getHeaders(status_code, len([]byte(body)), getMimeType("tmp.html"))
	conn.Write([]byte(headers))
	conn.Write([]byte(body))
}

func sendStaticFile(filepath string, conn net.Conn) {
	data, err := readFile(filepath)
	if err != nil {
		log.Printf("Unable to find file %s so did not respond\n", filepath)
		sendError(conn, NOT_FOUND_HTML, NOT_FOUND_CODE)
		return
	}
	headers := getHeaders(OK_CODE, len([]byte(data)), getMimeType(filepath))
	conn.Write([]byte(headers))
	conn.Write([]byte(data))
}
