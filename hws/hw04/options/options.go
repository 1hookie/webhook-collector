package options

import "fmt"

type Server struct {
	Host    string // дефолт "localhost"
	Port    int    // дефолт 8080
	Timeout int    // дефолт 30 (сек)
	Workers int    // дефолт 4
}

type Option func(*Server)

func NewServer(opts ...Option) Server {
	s := Server{
		Host:    "localhost",
		Port:    8080,
		Timeout: 30,
		Workers: 4,
	}
	for _, o := range opts {
		o(&s)
	}
	return s
}

func WithHost(host string) Option {
	return func(s *Server) { s.Host = host }
}

func WithPort(port int) Option {
	return func(s *Server) { s.Port = port }
}

func WithTimeout(seconds int) Option {
	return func(s *Server) { s.Timeout = seconds }
}

func WithWorkers(n int) Option {
	return func(s *Server) { s.Workers = n }
}

func (s Server) String() string {
	return fmt.Sprintf("%s:%d t=%d w=%d", s.Host, s.Port, s.Timeout, s.Workers)
}
