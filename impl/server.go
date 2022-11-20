package main

import (
	"context"
	"net"
)

type Server st

func (s *Server) Serve(l net.Listener) error {
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}

		s.setupConn(c)
		go s.ServeConn(context.Background(), c)
	}
}
func main() {
}
