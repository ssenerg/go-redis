package goredis

import (
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"
)

const defaultPort string = "6379"

type config struct {
	port string
}

func newConfig() config {
	var port string
	flag.StringVar(&port, "port", "", fmt.Sprintf("redis server port (default: %q)", defaultPort))
	flag.Parse()

	if port == "" {
		port = os.Getenv("PORT")
	}
	if port == "" {
		port = defaultPort
	}

	return config{
		port: port,
	}
}

type Server struct {
	config
	ln net.Listener
}


func NewServer() *Server {
	return &Server{
		config: newConfig(),
	}
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			slog.Error("error accepting connection", "error", err)
			continue
		}
		slog.Info("accepted connection", "remote_addr", conn.RemoteAddr())
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return err
	}
	slog.Info(fmt.Sprintf("server started and is listening on port %s", s.port))
	s.ln = ln
	s.acceptLoop()
	return nil
}