package youtrim

import (
	"github.com/mowshon/moviego"
)

type VideoTrimmerImpl struct{}

func (v *VideoTrimmerImpl) Trim(fullPath, trimmedPath string, start, end float64) error {
	first, _ := moviego.Load(fullPath)
	err := first.SubClip(start, end).Output(trimmedPath).Run()
	if err != nil {
		return err
	}
	return nil
}

func NewVideoTrimmerImpl() *VideoTrimmerImpl {
	return &VideoTrimmerImpl{}
}
