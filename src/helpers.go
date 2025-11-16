package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func readFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %v", path, err)
	}
	return string(data), nil
}

func getMimeType(path string) string {
	extension := strings.ToLower(filepath.Ext(path))
	switch extension {
	case ".html", ".htm":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".json":
		return "application/json"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".svg":
		return "image/svg+xml"
	case ".ico":
		return "image/x-icon"
	case ".txt":
		return "text/plain"
	default:
		return "application/octet-stream"
	}
}

func getHeaders(statusCode int, contentLength int, contentType string) string {
	var statusText string
	switch statusCode {
	case 200:
		statusText = "OK"
	case 400:
		statusText = "Bad Request"
	case 404:
		statusText = "Not Found"
	case 500:
		statusText = "Internal Server Error"
	default:
		statusText = "Unknown"
	}
	header := fmt.Sprintf(
		"HTTP/1.1 %d %s\r\nContent-Type: %s\r\nContent-Length: %d\r\nConnection: close\r\n\r\n",
		statusCode, statusText, contentType, contentLength,
	)

	return header
}
