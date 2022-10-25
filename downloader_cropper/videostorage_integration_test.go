package youtrim_test

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	youtrim "github.com/antunesleo/youtrim/downloader_cropper"
)

func TestVideoStorage(t *testing.T) {
	t.Run("test should save stream", func(t *testing.T) {
		filepath := "video.mp4"
		defer os.Remove(filepath)
		want := "zas"
		stringReader := strings.NewReader(want)
		stubStreamVideo := io.NopCloser(stringReader)

		vs := youtrim.NewVideoStorage()
		err := vs.CreateVideoFile(stubStreamVideo, filepath)

		assertNotError(err, t)

		videoFile, _ := ioutil.ReadFile(filepath)
		got := string(videoFile)

		if got != want {
			t.Errorf("want %s got %s", want, got)
		}
	})
}
