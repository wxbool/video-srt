## video-srt

这是一个可以识别视频语音自动生成字幕SRT文件的开源命令行工具。

本项目使用了阿里云的[OSS对象存储](https://www.aliyun.com/product/oss?spm=5176.12825654.eofdhaal5.13.e9392c4aGfj5vj&aly_as=K11FcpO8)、[录音文件识别](https://ai.aliyun.com/nls/filetrans?spm=5176.12061031.1228726.1.47fe3cb43I34mn)的相关业务接口。

Windows-GUI版本：[https://github.com/wxbool/video-srt-windows](https://github.com/wxbool/video-srt-windows)

## 下载安装
```shell
go get -u github.com/wxbool/video-srt
```

## 使用
###### 项目使用了 [ffmpeg](http://ffmpeg.org/) 依赖，请先下载安装，并设置环境变量.

* 设置服务接口配置（config.ini）
```ini
#字幕相关设置
[srt]
#智能分段处理：true（开启） false（关闭）
intelligent_block=true

#阿里云Oss对象服务配置
#文档：https://help.aliyun.com/document_detail/31827.html?spm=a2c4g.11186623.6.582.4e7858a85Dr5pA
[aliyunOss]
# OSS 对外服务的访问域名
endpoint=your.Endpoint
# 存储空间（Bucket）名称
bucketName=your.BucketName
# 存储空间（Bucket 域名）地址
bucketDomain=your.BucketDomain
accessKeyId=your.AccessKeyId
accessKeySecret=your.AccessKeySecret

#阿里云语音识别配置
#文档：
[aliyunClound]
# 在管控台中创建的项目Appkey，项目的唯一标识
appKey=your.AppKey
accessKeyId=your.AccessKeyId
accessKeySecret=your.AccessKeySecret
```

* 生成字幕文件（CLI）

```shell
go run main.go video.mp4
```

* 生成字幕文件（可执行文件 | [video-srt.exe](https://github.com/wxbool/video-srt/blob/master/video-srt.exe)）
```shell
video-srt video.mp4
```


## FAQ
* 支持哪些语言？
    * 视频字幕文本识别的核心服务是由阿里云`录音文件识别`业务提供的接口进行的，支持汉语普通话、方言、欧美英语等语言
* 如何才能使用这个工具？
    * 注册阿里云账号
    * 账号快速实名认证
    * 开通 `访问控制` 服务，并创建角色，设置开放 `OSS对象存储`、`智能语音交互` 的访问权限 
    * 开通 `OSS对象存储` 服务，并创建一个存储空间（Bucket）（读写权限设置为公共读）
    * 开通 `智能语音交互` 服务，并创建项目（根据使用场景选择识别语言以及偏好等）
    * 设置 `config.ini` 文件的配置项
    * 命令行执行（详见`使用`）