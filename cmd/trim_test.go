package cmd_test

import (
	"os"
	"testing"

	"github.com/antunesleo/youtrim/cmd"
	youtrim "github.com/antunesleo/youtrim/downloader_cropper"
	"github.com/stretchr/testify/assert"
)

func TestUnitTrimCmdRunner(t *testing.T) {
	type test struct {
		name string
		args []string
	}

	tests := []test{
		{name: "Should failed due to missing URL", args: []string{}},
		{name: "Should failed due to missing Start", args: []string{"--url", "url"}},
		{name: "Should failed due to missing End", args: []string{"--url", "url", "--start", "2"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			trimCmdRunner := cmd.NewTrimCmdRunner(&youtrim.StubTrimYtVideoUseCase{})

			command := cmd.BuildTrimCmd(*trimCmdRunner)
			command.SetArgs(test.args)
			err := command.Execute()
			assert.NotNil(t, err)
		})
	}
}

func TestIntegrationTrimCommand(t *testing.T) {
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
