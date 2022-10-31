package youtrim_test

import (
	"io"
	"strings"
	"testing"

	youtrim "github.com/antunesleo/youtrim/downloader_cropper"
	"github.com/stretchr/testify/assert"
)

func TestUnitTrimYtVideoUseCase(t *testing.T) {
	t.Run("Download and trim youtube video", func(t *testing.T) {
		stringReader := strings.NewReader("shiny!")
		stubStreamVideo := io.NopCloser(stringReader)

		stubYtDownloader := youtrim.StubYtDownloader{StubStream: stubStreamVideo}
		stubVideoStorage := youtrim.StubVideoStorage{}
		stubVideoTrimmer := youtrim.StubVideoTrimmer{}
		usecase := youtrim.NewTrimYtVideoUseCase(
			&stubYtDownloader,
			&stubVideoStorage,
			&stubVideoTrimmer,
		)

		url := "some-url"
		start := 3.0
		end := 6.0
		err := usecase.DownloadAndTrimYtVideo(url, start, end)

		assert.Nil(t, err)
		assert.Equal(t, 1, len(stubYtDownloader.Calls))
		assert.Equal(t, url, stubYtDownloader.Calls[0])
		assert.Equal(t, stubStreamVideo, stubVideoStorage.Calls[0].Stream)
		assert.Equal(t, youtrim.FullVideoPath, stubVideoStorage.Calls[0].Filepath)
		assert.Equal(t, 1, len(stubVideoTrimmer.Calls))
		assert.Equal(t, youtrim.FullVideoPath, stubVideoTrimmer.Calls[0].FullPath)
		assert.Equal(t, youtrim.TrimmedVideoPath, stubVideoTrimmer.Calls[0].TrimmedPath)
		assert.Equal(t, start, stubVideoTrimmer.Calls[0].Start)
		assert.Equal(t, end, stubVideoTrimmer.Calls[0].End)
	})

	t.Run("Download and trim youtube video failed due to download error", func(t *testing.T) {
		stringReader := strings.NewReader("shiny!")
		stubStreamVideo := io.NopCloser(stringReader)

		stubYtDownloader := youtrim.StubYtDownloader{StubStream: stubStreamVideo}
		stubVideoStorage := youtrim.StubVideoStorage{}
		stubVideoTrimmer := youtrim.StubVideoTrimmer{}
		usecase := youtrim.NewTrimYtVideoUseCase(
			&stubYtDownloader,
			&stubVideoStorage,
			&stubVideoTrimmer,
		)

		url := "please-fail"
		start := 3.0
		end := 6.0
		err := usecase.DownloadAndTrimYtVideo(url, start, end)

		assert.Equal(t, 1, len(stubYtDownloader.Calls))
		assert.Equal(t, url, stubYtDownloader.Calls[0])
		assert.Equal(t, 0, len(stubVideoStorage.Calls))
		assert.NotNil(t, err)
	})
}
