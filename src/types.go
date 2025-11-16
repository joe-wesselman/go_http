package main

type HttpRequest struct {
	Method  string
	Path    string
	Proto   string
	Headers map[string]string
}
