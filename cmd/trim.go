/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"strconv"

	youtrim "github.com/antunesleo/youtrim/downloader_cropper"
	"github.com/spf13/cobra"
)

type TrimYtVideoUseCaseInterface interface {
	DownloadAndTrimYtVideo(url string, start, end float64) error
}

type TrimCmdRunner struct {
	useCase TrimYtVideoUseCaseInterface
}

func (t *TrimCmdRunner) Run(cmd *cobra.Command, args []string) {
	urlFlag, _ := cmd.Flags().GetString("url")
	startFlag, _ := cmd.Flags().GetString("start")
	endFlag, _ := cmd.Flags().GetString("end")

	start, _ := strconv.ParseFloat(startFlag, 64)
	end, _ := strconv.ParseFloat(endFlag, 64)

	t.useCase.DownloadAndTrimYtVideo(urlFlag, float64(start), float64(end))
}

func init() {
	ytDownloader := youtrim.NewYtDownloader()
	videoStorage := youtrim.NewVideoStorage()
	videoTrimmer := youtrim.NewVideoTrimmerImpl()
	useCase := youtrim.NewTrimYtVideoUseCase(ytDownloader, videoStorage, videoTrimmer)
	trimCmdRunner := TrimCmdRunner{useCase: &useCase}

	var trimCmd = &cobra.Command{
		Use:   "trim",
		Short: "It downloads and trim a youtube video",
		Long: `It downloads and trim a youtube video.
	
	You should use the --url flag to inform the video that should be trimmed
	and the --start and --end flags to specify the seconds to trim`,
		Run: trimCmdRunner.Run,
	}
	rootCmd.AddCommand(trimCmd)
	trimCmd.PersistentFlags().String("url", "", "Youtube video URL to download and trim")
	trimCmd.PersistentFlags().String("start", "", "Video start (Seconds)")
	trimCmd.PersistentFlags().String("end", "", "Video end (Seconds)")
	trimCmd.MarkPersistentFlagRequired("url")
	trimCmd.MarkPersistentFlagRequired("start")
	trimCmd.MarkPersistentFlagRequired("end")
}
