package main

import (
	"flag"
	"fmt"
	"github.com.wxbool/video-srt/videosrt"
	"log"
	"os"
	"path/filepath"
	"time"
)

//定义配置文件
const CONFIG = "config.ini"

func main() {

	//致命错误捕获
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("")
			log.Printf("错误:\n%v", err)

			time.Sleep(time.Second * 5)
		}
	}()

	appDir, err := filepath.Abs(filepath.Dir(os.Args[0])) //应用执行根目录
	if err != nil {
		panic(err)
	}

	//初始化
	if len(os.Args) < 2 {
		os.Args = append(os.Args , "")
	}

	var video string

	//设置命令行参数
	flag.StringVar(&video, "f", "", "enter a video file waiting to be processed .")

	flag.Parse()

	if video == "" && os.Args[1] != "" && os.Args[1] != "-f" {
		video = os.Args[1]
	}

	//获取应用
	app := videosrt.NewApp(CONFIG)

	appDir = videosrt.WinDir(appDir)

	//初始化应用
	app.Init(appDir)

	//调起应用
	app.Run(videosrt.WinDir(video))

	//延迟退出
	time.Sleep(time.Second * 1)
}
