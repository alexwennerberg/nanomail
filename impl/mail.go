package main

import (
	"bufio"
	"io"
	"strings"
)

// Commands
const Fetch = "FETCH"
const Send = "SEND"
const GetKey = "GETKEY"

func ParseHeaders(mail io.Reader) map[string]string {
	var headers = make(map[string]string)
	scanner := bufio.NewScanner(mail)
	for scanner.Scan() {
		text := scanner.Text()
		b, a := strings.Cut(text, ": ")
		headers[b] = a
		// Break on new line
		if text == "" {
			continue
		}
	}
}
