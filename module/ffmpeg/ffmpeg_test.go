package ffmpeg

import (
	"fmt"
	"github.com/disintegration/imaging"
	"testing"
)

func TestGetCoverFromVideo(t *testing.T) {
	image, err := GetCoverFromVideo("./video.mp4", 10)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	img, err := imaging.Decode(image)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err = imaging.Save(img, "./cover.jpeg")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func TestAddWatermark(t *testing.T) {
	AddWatermark("./video.mp4", "./logo.png", "./waterVideo.mp4")
}
