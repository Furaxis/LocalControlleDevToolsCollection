package main

import (
	"github.com/Furaxis/LocalControlleDevToolsCollection/config"
	"github.com/Furaxis/LocalControlleDevToolsCollection/routes"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	println("passo 1 carregando...")
	routes.Init()
	println("passo 2 http://localhost:8080/api/v1/opening")
}

// url_base_teste : http://localhost:8080/ping
