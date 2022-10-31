package cmd_test

import (
	"os"
	"testing"

	"github.com/antunesleo/youtrim/cmd"
	youtrim "github.com/antunesleo/youtrim/downloader_cropper"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationTrim(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}

	ytDownloader := youtrim.NewYtDownloader()
	videoStorage := youtrim.NewVideoStorage()
	videoTrimmer := youtrim.NewVideoTrimmerImpl()
	useCase := youtrim.NewTrimYtVideoUseCase(ytDownloader, videoStorage, videoTrimmer)
	trimCmdRunner := cmd.NewTrimCmdRunner(&useCase)

	command := cmd.BuildTrimCmd(*trimCmdRunner)
	command.SetArgs(
		[]string{
			"--url",
			"https://www.youtube.com/watch?v=lyRt47FBpG8",
			"--start",
			"2",
			"--end",
			"4",
		},
	)

	err := command.Execute()
	trimmedPath := "./trimmedvideo.mp4"
	fullPath := "./video.mp4"
	trimmedVideo, trimmedErr := os.Open(trimmedPath)
	trimmedStat, _ := trimmedVideo.Stat()
	fullVideo, _ := os.Open(fullPath)
	fullStat, _ := fullVideo.Stat()
	assert.Nil(t, err)
	assert.Nil(t, trimmedErr)
	assert.True(t, trimmedStat.Size() < fullStat.Size())

	os.Remove(trimmedPath)
	os.Remove(fullPath)
}
