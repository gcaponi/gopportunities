package main

import (
	"fmt"

	"github.com/gcaponi/gopportunities/config"
	"github.com/gcaponi/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		logger.Errorf("Config Initialization error: %v", err)
	return
	}

	// Initialize Router
	router.Initialize()

}
