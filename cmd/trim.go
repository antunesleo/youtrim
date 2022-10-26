/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"strconv"
	"time"

	youtrim "github.com/antunesleo/youtrim/downloader_cropper"
	"github.com/spf13/cobra"
)

var trimCmd = &cobra.Command{
	Use:   "trim",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		urlFlag, _ := cmd.Flags().GetString("url")
		startFlag, _ := cmd.Flags().GetString("start")
		endFlag, _ := cmd.Flags().GetString("end")

		start, _ := strconv.ParseFloat(startFlag, 32)
		end, _ := strconv.ParseFloat(endFlag, 32)

		startDuration := time.Duration(start * float64(time.Second))
		endDuration := time.Duration(end * float64(time.Second))

		ytDownloader := youtrim.NewYtDownloader()
		videoStorage := youtrim.NewVideoStorage()
		videoTrimmer := youtrim.NewVideoTrimmerImpl()
		useCase := youtrim.NewTrimYtVideoUseCase(ytDownloader, videoStorage, videoTrimmer)
		useCase.DownloadAndTrimYtVideo(urlFlag, startDuration, endDuration)
	},
}

func init() {
	rootCmd.AddCommand(trimCmd)
	trimCmd.PersistentFlags().String("url", "", "Youtube video URL to download and trim")
	trimCmd.PersistentFlags().String("start", "", "Video start (Seconds)")
	trimCmd.PersistentFlags().String("end", "", "Video end (Seconds)")
	trimCmd.MarkPersistentFlagRequired("url")
	trimCmd.MarkPersistentFlagRequired("start")
	trimCmd.MarkPersistentFlagRequired("end")
}
