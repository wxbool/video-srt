package mylog

import (
	"log"
	"os"
)

//日志存储文件
const LOGFILE  = "log.txt"

//写入日志
func WriteLog(text ...interface{})  {
	if err := checkLogFile(LOGFILE); err != nil {
		panic(err.Error())
		return
	}
	logfile , err := os.OpenFile(LOGFILE , os.O_APPEND , os.ModePerm)
	if err != nil {
		panic(err.Error())
		return
	}

	defer logfile.Close() //关闭

	debugLog := log.New(logfile , "[info]" , log.Llongfile)
	debugLog.Println(text)
}


//检测日志文件
func checkLogFile (path string) (error) {
	file , err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			newfile , err := os.Create(path)
			if err != nil {
				return err
			}
			newfile.Close()
		}
	}
	file.Close()
	return nil
}

