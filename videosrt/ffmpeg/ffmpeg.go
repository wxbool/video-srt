package ffmpeg

import (
	"github.com/pkg/errors"
	"os/exec"
)

type Ffmpeg struct {
	Os string //ffmpeg 文件目录
}


//提取视频音频
func ExtractAudio (video string , tmpAudio string) (error) {
	ts := exec.Command("ffmpeg" , "-version")
	if _ , err := ts.CombinedOutput() ; err != nil {
		return errors.New("请先安装 ffmpeg 依赖 ，并设置环境变量")
	}

	cmd := exec.Command("ffmpeg" , "-i" , video , "-ar" , "16000" , tmpAudio)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
