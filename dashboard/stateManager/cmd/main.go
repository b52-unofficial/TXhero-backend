package main

import (
	"github.com/b52-unofficial/TXhero-backend/config"
	"github.com/b52-unofficial/TXhero-backend/dashboard/stateManager/common/logger"
	"github.com/b52-unofficial/TXhero-backend/dashboard/stateManager/worker"
	"sync"

	_ "github.com/b52-unofficial/TXhero-backend/docs"
)

func main() {
	// DB connection
	//db.ConnectDB()
	cfg := config.GetConfig()
	logger.SetLogLevel(cfg.LogLevel)

	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)
	go RunWorker()
}

func RunWorker() {
	worker.RegisterCron()
}
