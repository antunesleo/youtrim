package youtrim_test

import (
	"io"
	"strings"
	"testing"
	"time"

	"github.com/antunesleo/youtrim"
)

func TestTrimYtVideoUseCase(t *testing.T) {
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
		start := 3 * time.Second
		end := 6 * time.Second
		err := usecase.DownloadAndTrimYtVideo(url, start, end)

		assertNotError(err, t)
		assertCallCount(1, len(stubYtDownloader.Calls), t)
		assertCalledWith(url, stubYtDownloader.Calls[0], t)

		assertCallCount(1, len(stubVideoStorage.Calls), t)
		if stubStreamVideo != stubVideoStorage.Calls[0].Stream {
			t.Errorf("Want %d got %d", stubStreamVideo, stubVideoStorage.Calls[0].Stream)
		}

		if youtrim.FullVideoPath != stubVideoStorage.Calls[0].Filepath {
			t.Errorf("Want %s got %s", youtrim.FullVideoPath, stubVideoStorage.Calls[0].Filepath)
		}

		assertCallCount(1, len(stubVideoTrimmer.Calls), t)

		if youtrim.FullVideoPath != stubVideoTrimmer.Calls[0].FullPath {
			t.Errorf(
				"Want %s got %s",
				youtrim.FullVideoPath,
				stubVideoTrimmer.Calls[0].FullPath,
			)
		}

		if youtrim.TrimmedVideoPath != stubVideoTrimmer.Calls[0].TrimmedPath {
			t.Errorf(
				"Want %s got %s",
				youtrim.TrimmedVideoPath,
				stubVideoTrimmer.Calls[0].TrimmedPath,
			)
		}

		if start != stubVideoTrimmer.Calls[0].Start {
			t.Errorf("Want %d got %d", start, stubVideoTrimmer.Calls[0].Start)
		}

		if end != stubVideoTrimmer.Calls[0].End {
			t.Errorf("Want %d got %d", end, stubVideoTrimmer.Calls[0].End)
		}

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
		start := 3 * time.Second
		end := 6 * time.Second
		err := usecase.DownloadAndTrimYtVideo(url, start, end)

		assertCallCount(1, len(stubYtDownloader.Calls), t)
		assertCalledWith(url, stubYtDownloader.Calls[0], t)
		assertCallCount(0, len(stubVideoStorage.Calls), t)

		if err == nil {
			t.Errorf("want error to be nil got %d", err)
		}
	})
}

func assertCalledWith(want, got string, t *testing.T) {
	if want != got {
		t.Errorf("want URL %s got %s", want, got)
	}
}

func assertCallCount(want, got int, t *testing.T) {
	if want != got {
		t.Fatalf("want number calls %d got %d", want, got)
	}
}
