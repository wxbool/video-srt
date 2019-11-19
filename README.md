## video-srt

这是一个视频自动生成字幕SRT文件的开源解决方案。

本项目使用了阿里云的[OSS对象存储](https://www.aliyun.com/product/oss?spm=5176.12825654.eofdhaal5.13.e9392c4aGfj5vj&aly_as=K11FcpO8)、[录音文件识别](https://ai.aliyun.com/nls/filetrans?spm=5176.12061031.1228726.1.47fe3cb43I34mn)的相关业务接口。


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