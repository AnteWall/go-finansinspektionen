package main

import (
	"github.com/joho/godotenv"
	"go-finansinspektionen/internal/insider"
	"go-finansinspektionen/pkg/utils"
)

func main() {
	_ = godotenv.Load()
	logger := utils.NewLogger()
	logger.Info("starting scraper")

	client := insider.NewClient(insider.WithDebug(true))

	err := client.GetLatestTrades()
	if err != nil {
		logger.Error(err)
	}
}
