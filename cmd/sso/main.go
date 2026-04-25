package main

import (
	"os"

	"github.com/bagusyanuar/go-simrs/internal/shared/bootstrap"
	"github.com/bagusyanuar/go-simrs/internal/shared/config"
	"github.com/bagusyanuar/go-simrs/internal/shared/container"
)

func main() {
	// 1. Load Config
	conf := config.LoadConfig()

	// 2. Override Port for SSO if needed (Default 8081 for SSO)
	if os.Getenv("APP_PORT") == "" {
		conf.AppPort = "8081"
	}

	// 3. Initialize Logger
	config.InitLogger(conf)

	// 4. Initialize Database
	db := config.InitDB(conf)

	// 5. Initialize Dependency Container
	deps := container.NewContainer(db, conf)

	// 6. Start Server
	bootstrap.Start(conf, deps)
}
