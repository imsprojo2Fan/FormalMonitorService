
1>官方文档环境安装说明
如果HCNetSDKCom目录以及libhcnetsdk.so、libhpr.so、libHCCore文件和可执行文件在同一级目录下，则使用同级目录下的库文件;
如果不在同一级目录下，则需要将以上文件的目录加载到动态库搜索路径中，设置的方式有以下几种:
一.	将网络SDK各动态库路径加入到LD_LIBRARY_PATH环境变量
	1.在终端输入：export  LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/XXX:/XXX/HCNetSDKCom      只在当前终端起作用
	2. 修改~/.bashrc或~/.bash_profile，最后一行添加 export  LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/XXX:/XXX/HCNetSDKCom，保存之后，使用source  .bashrc执行该文件 ，当前用户生效
	3. 修改/etc/profile，添加内容如第2条，同样保存之后使用source执行该文件  所有用户生效

二．在/etc/ld.so.conf文件结尾添加网络sdk库的路径，如/XXX和/XXX/HCNetSDKCom/，保存之后，然后执行ldconfig 

三．可以将网络sdk各依赖库放入到/lib64或usr/lib64下

四．可以在Makefile中使用-Wl,-rpath来指定动态路径，但是需要将网络sdk各个动态库都用 –l方式显示加载进来
	比如：-Wl,-rpath=/XXX:/XXX/HCNetSDKCom -lhcnetsdk  -lhpr –lHCCore –lHCCoreDevCfg –lStreamTransClient –lSystemTransform –lHCPreview –lHCAlarm –lHCGeneralCfgMgr –lHCIndustry –lHCPlayBack –lHCVoiceTalk –lanalyzedata -lHCDisplay


推荐使用一或二的方式，但要注意优先使用的是同级目录下的库文件

2>二次开发使用手册

	1.软件环境要求:linux Ubuntu 64位系统、gcc编译环境、golang 64位、安装ffmpeg(用于视频格式转换)
	2.程序入口为main.go、build.sh可编译cgo
	3.\FormalMonitorService\golangMonitor\linux64\proj
	4.Capture.c可更改摄像机参数信息，如有更改需重新运行build.sh
	5.1>go get "github.com/robfig/cron"
	5.2>go get "github.com/lucky2me/log"
	5.3>go日志文件指定目录为"/opt/go/FormalMonitorService/golangMonitor/linux64/proj/goLog/"如有变动请更改
	6.具体详情请参看代码
	7.CentOS环境下"go build command-line-arguments: invalid flag in #cgo LDFLAGS: -Wl,-rpath=./:./HCNetSDKCom:../lib"
	  解决:export CGO_CXXFLAGS_ALLOW=".*" 
		   export CGO_LDFLAGS_ALLOW=".*" 
		   export CGO_CFLAGS_ALLOW=".*" 
		   (可将该环境变量设置到全局环境变量：vim /etc/profile 将export添加到文件末尾保存，重新打开会话窗口即可)
	8.更改了CapPicture.c里的方法名或参数时需更改CapPicture.h文件里的内容方可奏效！！！！！
	9.进入main.go文件夹 "nohup go run main.go >/dev/null 2>&1 &" 启动程序