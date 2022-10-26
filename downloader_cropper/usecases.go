package youtrim

import (
	"fmt"
	"io"
)

const FullVideoPath = "video.mp4"
const TrimmedVideoPath = "trimmedvideo.mp4"

type YtDownloader interface {
	DownloadYtVideo(URL string) (io.ReadCloser, error)
}

type VideoStorage interface {
	CreateVideoFile(stream io.ReadCloser, filepath string) error
}

type VideoTrimmer interface {
	Trim(fullPath, trimmedPath string, start, end float64) error
}

type TrimYtVideoUseCase struct {
	ytDownloader YtDownloader
	videoStorage VideoStorage
	videoTrimmer VideoTrimmer
}

func (t *TrimYtVideoUseCase) DownloadAndTrimYtVideo(url string, start, end float64) error {
	stream, err := t.ytDownloader.DownloadYtVideo(url)
	if err != nil {
		return fmt.Errorf("Failed to download and trim video; %w", err)
	}
	t.videoStorage.CreateVideoFile(stream, FullVideoPath)
	t.videoTrimmer.Trim(FullVideoPath, TrimmedVideoPath, start, end)
	return nil
}

func NewTrimYtVideoUseCase(yd YtDownloader, vs VideoStorage, vt VideoTrimmer) TrimYtVideoUseCase {
	return TrimYtVideoUseCase{ytDownloader: yd, videoStorage: vs, videoTrimmer: vt}
}
