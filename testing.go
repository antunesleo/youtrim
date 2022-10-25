package youtrim

import (
	"io"
	"time"
)

type StubYtDownloader struct {
	StubStream io.ReadCloser
	Calls      []string
}

func (s *StubYtDownloader) DownloadYtVideo(URL string) (io.ReadCloser, error) {
	s.Calls = append(s.Calls, URL)
	if URL == "please-fail" {
		return nil, ErrFailedToDownloadVideo
	}
	return s.StubStream, nil
}

type CreateVideoFileArgs struct {
	Stream   io.ReadCloser
	Filepath string
}

type StubVideoStorage struct {
	Calls []CreateVideoFileArgs
}

func (s *StubVideoStorage) CreateVideoFile(stream io.ReadCloser, filepath string) error {
	s.Calls = append(s.Calls, CreateVideoFileArgs{stream, filepath})
	return nil
}

type VideoTrimmerTrimArgs struct {
	FullPath, TrimmedPath string
	Start, End            time.Duration
}

type StubVideoTrimmer struct {
	Calls []VideoTrimmerTrimArgs
}

func (s *StubVideoTrimmer) Trim(fullPath, trimmedPath string, start, end time.Duration) error {
	s.Calls = append(s.Calls, VideoTrimmerTrimArgs{fullPath, trimmedPath, start, end})
	return nil
}
