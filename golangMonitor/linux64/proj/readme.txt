接口说明
-----------------------------------------------------------------------------开始
自动截图接口
必填参数：channelNum<int>
返回结果：
	  1>成功：20180613-095417.jpeg
	  2>失败：[AutoSnap]|Err:7|msg:1528855254  -->方法类型、错误码\文件名时间戳

http://192.168.213.133/autoSnap?channelNum=1

/*按时间截图接口
必填参数：channelNum<int> curTime<long>
返回结果：
	  1>成功：20180613-095417.jpeg
	  2>失败：[AutoSnap]|Err:7|msg:1528855254  -->方法类型、错误码\文件名时间戳

http://192.168.213.133/snapByTime?channelNum=1&curTime=1528955647

自动下载接口
必填参数：channelNum<int> 可选参数timeCount<long> (例:300即下载的视频为当前时间至前300秒内的视频)
返回结果：
	  1>成功：20180613-094820.mp4
	  2>失败：[AutoDownload]|Err:7|msg:1528855254  -->方法类型、错误码\文件名时间戳

http://192.168.213.133/autoDownload?channelNum=1

按时间下载接口
必填参数：channelNum<int> startTime<long> endTime<long>
返回结果：
	  1>成功：20180613-094820.mp4
	  2>失败：[DownloadByTime]|Err:7|msg:1528855254  -->方法类型、错误码\文件名时间戳

http://192.168.213.133/downloadByTime?channelNum=1&startTime=1528791412&endTime=1528791712

获取图片列表接口
http://192.168.213.133/picList/

获取视频列表接口
http://192.168.213.133/videoList/

注:如自动下载视频失败时，可通过按时间下载接口再次下载

接口说明
-----------------------------------------------------------------------------结束