package main

import (
	"flag"
	"log"

	"github.com/0x5w4/kredit-plus/writer-service/config"
	"github.com/0x5w4/kredit-plus/writer-service/internal/server"
)

func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Failed to init config: %v", err)
	}

	s := server.NewServer(cfg)
	if err := s.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
