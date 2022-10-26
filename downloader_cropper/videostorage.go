package youtrim

import (
	"errors"
	"io"
	"os"
)

var ErrFailedToSaveVideoStream = errors.New("Failed to save video stream")

type VideoStorageImpl struct{}

func (vs *VideoStorageImpl) CreateVideoFile(stream io.ReadCloser, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return ErrFailedToSaveVideoStream
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return ErrFailedToSaveVideoStream
	}
	return nil
}

func NewVideoStorage() *VideoStorageImpl {
	return &VideoStorageImpl{}
}
