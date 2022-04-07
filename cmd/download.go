/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-finansinspektionen/internal/insider"
	"go-finansinspektionen/pkg/utils"
	"go.uber.org/zap"
	"time"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Downloads historical insider data",
	Long: `Downloads historical insider data from Finansinspektions insynsregister. 
It will download data split daily for each day. Todays date will be all available data up until the execution time.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := utils.NewLogger()
		startTime, err := time.Parse("2006-01-02", StrStartTime)
		if err != nil {
			return err
		}
		endTime := time.Now()
		if StrEndTime != "" {
			endTime, err = time.Parse("2006-01-02", StrEndTime)
			if err != nil {
				return err
			}
		}
		client := insider.NewClient(insider.WithDebug(Debug))

		tomorrow := time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, endTime.Location()).AddDate(0, 0, 1)
		for curr := startTime; curr.Before(tomorrow); curr = curr.AddDate(0, 0, 1) {
			output := fmt.Sprintf("%s/%s.csv", Output, curr.Format("2006-01-02"))
			logger.With(zap.Time("date", curr),zap.String("output", output)).Info("downloading file")
			t, err := client.GetTransactions(curr)
			if err != nil {
				logger.Fatal(err)
			}
			err = utils.SaveAsCSV(output, t)
			if err != nil {
				logger.Error(err)
			}
		}
		return nil
	},
}

var (
	Output       string
	StrStartTime string
	StrEndTime   string
	Debug bool
)

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVarP(&Output, "output", "o", "", "Output folder")
	downloadCmd.MarkFlagRequired("output")
	downloadCmd.Flags().StringVarP(&StrStartTime, "startdate", "s", "", "Start date")
	downloadCmd.MarkFlagRequired("startdate")
	downloadCmd.Flags().StringVarP(&StrEndTime, "enddate", "e", "", "End date")
	downloadCmd.Flags().BoolVarP(&Debug, "debug", "d", false, "debug http client")
}
