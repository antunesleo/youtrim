package youtrim_test

import (
	"testing"

	youtrim "github.com/antunesleo/youtrim/downloader_cropper"
)

const testVideoURL = "https://www.youtube.com/watch?v=jhFDyDgMVUI"
const invalidTestVideoURL = "https://www.youtube.com/watch?v=BrokenLink"

func TestDownload(t *testing.T) {
	t.Run("should download video", func(t *testing.T) {
		downloader := youtrim.NewYtDownloader()
		stream, err := downloader.DownloadYtVideo(testVideoURL)
		assertNotError(err, t)
		stream.Close()
	})

	t.Run("should failed to download when broken link", func(t *testing.T) {
		downloader := youtrim.NewYtDownloader()
		_, err := downloader.DownloadYtVideo(invalidTestVideoURL)
		assertError(err, youtrim.ErrFailedToDownloadVideo, t)
	})
}

func assertNotError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Want error to be nil, got %d", err)
	}
}

func assertError(want error, got error, t *testing.T) {
	if want != got {
		t.Errorf("Want %d, got %d", want, got)
	}
}
