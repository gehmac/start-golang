package main

import (
	"github.com/gehmac/anime-list-golang/config"
	"github.com/gehmac/anime-list-golang/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
