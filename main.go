package main

import (
	"github.com/Furaxis/LocalControlleDevToolsCollection/routes"
)

func main() {
	print("http://localhost:8080/api/v1/opening")
	routes.Init()
}

// url_base_teste : http://localhost:8080/ping
