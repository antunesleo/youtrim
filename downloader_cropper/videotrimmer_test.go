package youtrim_test

import (
	"os"
	"testing"

	youtrim "github.com/antunesleo/youtrim/downloader_cropper"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationVideoTrimmerImpl(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	fullPath := "./../videos/testvideo.mp4"
	trimmedPath := "./trimmed.mp4"

	videoTrimmer := youtrim.NewVideoTrimmerImpl()
	err := videoTrimmer.Trim(fullPath, trimmedPath, 2.0, 5.0)

	assert.Nil(t, err)

	trimmedVideo, trimmedErr := os.Open(trimmedPath)
	assert.Nil(t, trimmedErr)

	trimmedStat, _ := trimmedVideo.Stat()
	fullVideo, _ := os.Open(fullPath)
	fullStat, _ := fullVideo.Stat()

	assert.False(t, trimmedStat.Size() > fullStat.Size())
	os.Remove(trimmedPath)
}
