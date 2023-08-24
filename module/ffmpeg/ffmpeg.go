package ffmpeg

import (
	"bytes"
	"fmt"
	"github.com/Biu-X/TikTok/module/log"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
)

func GetCoverFromVideoWithDefault(path string) (io.Reader, error) {
	return GetCoverFromVideo(path, 1)
}

func GetCoverFromVideo(inFileName string, frameNum int) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf).
		Run()
	if err != nil {
		log.Logger.Errorf("GetCoverFromVideo Error: %v\n", err)
		return nil, err
	}
	return buf, nil
}

// AddWatermark 有个问题,添加水印后视频没有声音了
func AddWatermark(inVideoPath, watermarkPath, outVideoPath string) {
	// show watermark with size 64:-1 in the top left corner after seconds 1
	overlay := ffmpeg.Input(watermarkPath).Filter("scale", ffmpeg.Args{"20:-1"})
	err := ffmpeg.Filter(
		[]*ffmpeg.Stream{
			ffmpeg.Input(inVideoPath),
			overlay,
		}, "overlay", ffmpeg.Args{"10:10"}, ffmpeg.KwArgs{"enable": "gte(t,0)"}).
		Output(outVideoPath).OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
