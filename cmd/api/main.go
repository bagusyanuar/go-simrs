package main

import (
	"github.com/bagusyanuar/go-simrs/internal/shared/bootstrap"
	"github.com/bagusyanuar/go-simrs/internal/shared/config"
	"github.com/bagusyanuar/go-simrs/internal/shared/container"
)

func main() {
	// 1. Load Config
	conf := config.LoadConfig()

	// 2. Initialize Logger
	config.InitLogger(conf)

	// 3. Initialize Database
	db := config.InitDB(conf)

	// 4. Initialize Dependency Container
	deps := container.NewContainer(db, conf)

	// 5. Start Server
	bootstrap.Start(conf, deps)
}
