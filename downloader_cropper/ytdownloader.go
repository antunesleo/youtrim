package youtrim

import (
	"errors"
	"io"
	"strings"

	"github.com/kkdai/youtube/v2"
)

var ErrFailedToDownloadVideo = errors.New("Failed to download video")

type YtDownloaderImpl struct{}

func NewYtDownloader() *YtDownloaderImpl {
	return &YtDownloaderImpl{}
}

func (y *YtDownloaderImpl) DownloadYtVideo(URL string) (io.ReadCloser, error) {
	videoId := y.getVideoId(URL)
	client := youtube.Client{}

	video, err := client.GetVideo(videoId)
	if err != nil {
		return nil, ErrFailedToDownloadVideo
	}

	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		return nil, ErrFailedToDownloadVideo
	}
	return stream, err
}

func (*YtDownloaderImpl) getVideoId(URL string) string {
	splitUrl := strings.Split(URL, "watch?v=")
	videoId := splitUrl[len(splitUrl)-1]
	return videoId
}
