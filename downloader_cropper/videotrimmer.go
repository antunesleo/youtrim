package youtrim

import (
	"time"

	"github.com/jtguibas/cinema"
)

type VideoTrimmerImpl struct{}

func (v *VideoTrimmerImpl) Trim(fullPath, trimmedPath string, start, end time.Duration) error {
	fullVideo, _ := cinema.Load(fullPath)
	fullVideo.Trim(start*time.Second, end*time.Second)
	fullVideo.Render(trimmedPath)
	return nil
}

func NewVideoTrimmerImpl() *VideoTrimmerImpl {
	return &VideoTrimmerImpl{}
}
