package youtrim_test

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	youtrim "github.com/antunesleo/youtrim/downloader_cropper"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationVideoStorage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	t.Run("test should save stream", func(t *testing.T) {
		filepath := "video.mp4"
		defer os.Remove(filepath)
		want := "zas"
		stringReader := strings.NewReader(want)
		stubStreamVideo := io.NopCloser(stringReader)

		vs := youtrim.NewVideoStorage()
		err := vs.CreateVideoFile(stubStreamVideo, filepath)

		assert.Nil(t, err)

		videoFile, _ := ioutil.ReadFile(filepath)
		got := string(videoFile)

		assert.Equal(t, got, want)
	})
}
