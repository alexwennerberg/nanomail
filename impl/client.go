package main

import (
	"fmt"
	"strings"
	"time"
)

type Nanomail struct {
	Signature string
	From      string
	To        string
	SentAt    time.Time
	ThreadId  string
	Subject   string
	Body      string
}

func (n Nanomail) Validate() error {
	for _, v := range []string{n.Signature, n.From, n.To, n.Subject} {
		if strings.Contains(v, "\n") {
			return fmt.Errorf("Field cannot have newline")
		}
	}
	return nil
}

func (n *Nanomail) Sign() {
	n.Signatrue = ""
}

func (n Nanomail) StringMinusSignature() string {
	return fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s",
		n.From, n.To, n.Subject, n.Body)
}
func (n Nanomail) String() string {
	return fmt.Sprintf("Signature :%s\n%s",
		n.Signature, n.StringMinusSignature())
}

func main() {
	// usage: ./this sigfile from@example.com  dest@example.com
	// Reads message from stdin
}
