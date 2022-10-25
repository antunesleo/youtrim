package main

import (
	youtrim "github.com/antunesleo/youtrim"
)

func main() {
	url := "https://www.youtube.com/watch?v=kJ8e7-hFg-U"
	ytDownloader := youtrim.NewYtDownloader()
	videoStorage := youtrim.NewVideoStorage()
	videoTrimmer := youtrim.NewVideoTrimmerImpl()
	trimYtVideoUseCase := youtrim.NewTrimYtVideoUseCase(ytDownloader, videoStorage, videoTrimmer)
	trimYtVideoUseCase.DownloadAndTrimYtVideo(url, 2.0, 3.0)
}
