package youtrim_test

import (
	"testing"

	youtrim "github.com/antunesleo/youtrim/downloader_cropper"
	"github.com/stretchr/testify/assert"
)

const testVideoURL = "https://www.youtube.com/watch?v=jhFDyDgMVUI"
const invalidTestVideoURL = "https://www.youtube.com/watch?v=BrokenLink"

func TestIntegrationDownloader(t *testing.T) {
	t.Run("should download video", func(t *testing.T) {
		downloader := youtrim.NewYtDownloader()
		stream, err := downloader.DownloadYtVideo(testVideoURL)
		assert.Nil(t, err)
		stream.Close()
	})

	t.Run("should failed to download when broken link", func(t *testing.T) {
		downloader := youtrim.NewYtDownloader()
		_, err := downloader.DownloadYtVideo(invalidTestVideoURL)
		assert.Equal(t, err, youtrim.ErrFailedToDownloadVideo)
	})
}
