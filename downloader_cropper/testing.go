package youtrim

import (
	"io"
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
	Start, End            float64
}

type StubVideoTrimmer struct {
	Calls []VideoTrimmerTrimArgs
}

func (s *StubVideoTrimmer) Trim(fullPath, trimmedPath string, start, end float64) error {
	s.Calls = append(s.Calls, VideoTrimmerTrimArgs{fullPath, trimmedPath, start, end})
	return nil
}

type DownloadAndTrim struct {
	url        string
	start, end float64
}

type StubTrimYtVideoUseCase struct {
	StubStream io.ReadCloser
	Calls      []DownloadAndTrim
}

func (s *StubTrimYtVideoUseCase) DownloadAndTrimYtVideo(url string, start, end float64) error {
	s.Calls = append(s.Calls, DownloadAndTrim{url: url, start: start, end: end})
	return nil
}
