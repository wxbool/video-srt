package videosrt

import (
	"os"
)

//校验文件是否存在
func VaildVideo (video string) bool {
	_, err := os.Stat(video)  //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
