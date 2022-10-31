package cmd_test

import (
	"os"
	"testing"

	"github.com/antunesleo/youtrim/cmd"
	youtrim "github.com/antunesleo/youtrim/downloader_cropper"
)

func TestRun(t *testing.T) {
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

	trimmedPath := "./trimmedvideo.mp4"
	fullPath := "./video.mp4"

	err := command.Execute()
	if err != nil {
		t.Errorf("expected error to be nil")
	}

	trimmedVideo, trimmedErr := os.Open(trimmedPath)
	if trimmedErr != nil {
		t.Fatalf("failed to opened trimmed file %d", trimmedErr)
	}
	trimmedStat, _ := trimmedVideo.Stat()

	fullVideo, _ := os.Open(fullPath)
	fullStat, _ := fullVideo.Stat()

	if trimmedStat.Size() > fullStat.Size() {
		t.Errorf("Expected trimmed video size to be smaller than full video")
	}
}
