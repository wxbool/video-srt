package videosrt

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

//日志输出
func Log(agrs ...interface{}) {
	fmt.Println(agrs ...)
}

//Windows下Dir路径转换
func WinDir(dir string) string {
	return strings.Replace(dir , "\\" , "/" , -1)
}

//获取文件名称（不带后缀）
func GetFileBaseName(filepath string) string {
	basefile := path.Base(filepath)
	ext := path.Ext(filepath)

	return strings.Replace(basefile , ext , "" , 1)
}

//检验目录是否存在
func DirExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}


//创建目录
func CreateDir(path string , all bool) error {
	var err error
	if all {
		err = os.MkdirAll(path, os.ModePerm)
	} else {
		err = os.Mkdir(path, os.ModePerm)
	}
	if err != nil {
		return err
	}
	return nil
}

//获取随机字符串
func GetRandomCodeString(len int) string {
	rand.Seed(time.Now().Unix())  //设置随机种子

	seed := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seedArr := strings.Split(seed , "")

	result := []string{}
	index := 0
	for index < len {
		s := GetIntRandomNumber(0 , 61)
		result = append(result , seedArr[s])

		index++
	}

	return strings.Join(result , "")
}


//获取某范围的随机整数
func GetIntRandomNumber(min int64 , max int64) int64 {
	return rand.Int63n(max - min) + min
}


//字幕时间戳转换
func SubtitleTimeMillisecond(time int64) string {
	var miao int64 = 0
	var min int64 = 0
	var hours int64 = 0
	var millisecond int64 = 0

	millisecond = (time % 1000)
	miao = (time / 1000)

	if miao > 59 {
		min = (time / 1000) / 60
		miao = miao % 60
	}
	if min > 59 {
		hours = (time / 1000) / 3600
		min = min % 60
	}

	//00:00:06,770
	var miaoText = RepeatStr(strconv.FormatInt(miao , 10) , "0" , 2 , true)
	var minText = RepeatStr(strconv.FormatInt(min , 10) , "0" , 2 , true)
	var hoursText = RepeatStr(strconv.FormatInt(hours , 10) , "0" , 2 , true)
	var millisecondText = RepeatStr(strconv.FormatInt(millisecond , 10) , "0" , 3 , true)

	return hoursText + ":" + minText + ":" + miaoText + "," + millisecondText
}


func RepeatStr(str string , s string , length int , before bool) string {
	ln := len(str)

	if ln >= length {
		return str
	}

	if before {
		return  strings.Repeat(s , (length - ln)) + str
	} else {
		return  str + strings.Repeat(s , (length - ln))
	}
}


//打印对象转JSON数据
func DumpObjectToJson(data interface{} , title string) {
	if data != nil {
		jsonData , _ := json.Marshal(data)
		fmt.Println(title , string(jsonData))
	}
}