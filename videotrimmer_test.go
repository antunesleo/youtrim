package youtrim_test

import (
	"os"
	"testing"

	"github.com/antunesleo/youtrim"
)

func TestVideoTrimmerImpl(t *testing.T) {
	fullPath := "./videos/testvideo.mp4"
	trimmedPath := "./trimmed.mp4"

	videoTrimmer := youtrim.NewVideoTrimmerImpl()
	err := videoTrimmer.Trim(fullPath, trimmedPath, 2.0, 5.0)

	assertNotError(err, t)

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
