package ffmpeg

import (
	"os/exec"

	"github.com/Biu-X/TikTok/module/log"
)

// CoverSnap 使用 cv2 的方式截取视频封面，做了一些特殊处理，pip install coversnap 即可（pip默认添加到环境变量）
func CoverSnap(inputVideoPath, outputCoverPath string) {
	cmdName := "coversnap"
	cmdArgs := []string{"-i", inputVideoPath, "-o", outputCoverPath}
	// 创建命令对象
	cmd := exec.Command(cmdName, cmdArgs...)
	err := cmd.Run()
	if err != nil {
		log.Logger.Errorf("CoverSnap Error: %v", err)
	}
	log.Logger.Info("CoverSnap Success")
}
