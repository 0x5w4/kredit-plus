package server

import "github.com/0x5w4/kredit-plus/config"

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s Server) Run() error {
	return nil
}

func (s Server) RunSwagDoc() error {
	return nil
}

func (s Server) RunRpc() error {
	return nil
}
