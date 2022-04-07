package main

import (
	"github.com/AnteWall/go-finansinspektionen/internal/insider"
	"github.com/AnteWall/go-finansinspektionen/pkg/utils"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	logger := utils.NewLogger()
	logger.Info("starting scraper")

	client := insider.NewClient(insider.WithDebug(true))

	t, err := client.GetTodayTransactions()
	if err != nil {
		logger.Fatal(err)
	}

	err = utils.SaveAsCSV("output/test-latest.csv", t)
	if err != nil {
		logger.Error(err)
	}
}
