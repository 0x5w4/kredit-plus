package main

import (
	"flag"
	"log"

	"github.com/0x5w4/kredit-plus/api-gateway-service/config"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/server"
)

//	@title			API Gateway Kredit Plus
//	@version		1.0
//	@description	This is a sample API on Kredit Plus Microservice.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:5001
//	@BasePath	/api/v1

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
