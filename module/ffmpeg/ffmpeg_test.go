package ffmpeg

import (
	"testing"
)

func TestGetCoverFromVideo(t *testing.T) {
}

func TestAddWatermark(t *testing.T) {
	AddWatermark("./video.mp4", "./logo.png", "./waterVideo.mp4")
}
