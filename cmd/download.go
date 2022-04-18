package cmd

import (
	"fmt"
	"github.com/AnteWall/go-finansinspektionen/pkg/insider"
	"github.com/AnteWall/go-finansinspektionen/pkg/utils"
	"github.com/spf13/cobra"

	"go.uber.org/zap"
	"time"
)

// DownloadCmd represents the download command
var DownloadCmd = &cobra.Command{
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
			logger.With(zap.Time("date", curr), zap.String("output", output)).Info("downloading file")
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
	Debug        bool
)

func init() {
	RootCmd.AddCommand(DownloadCmd)

	DownloadCmd.Flags().StringVarP(&Output, "output", "o", "", "Output folder")
	DownloadCmd.MarkFlagRequired("output")
	DownloadCmd.Flags().StringVarP(&StrStartTime, "startdate", "s", "", "Start date")
	DownloadCmd.MarkFlagRequired("startdate")
	DownloadCmd.Flags().StringVarP(&StrEndTime, "enddate", "e", "", "End date")
	DownloadCmd.Flags().BoolVarP(&Debug, "debug", "d", false, "debug http client")
}
