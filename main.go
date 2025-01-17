package main

import (
	"job-openings/config"
	"job-openings/router"
)

func main() {
	logger := config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.ErrorF("Error at init config: %v", err)
	}

	router.InitRouter()
}
