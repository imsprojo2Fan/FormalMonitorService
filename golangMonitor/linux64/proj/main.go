package main

import (
	"fmt"
	"net/http"
	"strings"
	//"log"
	"time"
	"strconv"
	"os/exec"
	"os"
	"io/ioutil"
	//"path/filepath"
	//"io/ioutil"
	"github.com/lucky2me/log"
	"reflect"

)

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -L ../lib -Wl,-rpath=./:./HCNetSDKCom:../lib   -lCapPicture -lhcnetsdk
#include "../../include/Sdk.h"
#include "CapPicture.h"
*/
import "C"


/*
接口说明
-----------------------------------------------------------------------------开始
自动截图接口
必填参数：channelNum<int>
返回结果：
	  1>成功：20180613-095417.jpeg
	  2>失败：[AutoSnap]|Err:7|msg:1528855254  -->方法类型、错误码、文件名时间戳

http://192.168.213.133/autoSnap?channelNum=1

/*按时间截图接口
必填参数：channelNum<int> curTime<long>
返回结果：
	  1>成功：20180613-095417.jpeg
	  2>失败：[AutoSnap]|Err:7|msg:1528855254  -->方法类型、错误码、文件名时间戳

http://192.168.213.133/snapByTime?channelNum=1&curTime=1528955647

自动下载接口
必填参数：channelNum<int> 可选参数timeCount<long> (例:300即下载的视频为当前时间至前300秒内的视频)
返回结果：
	  1>成功：20180613-094820.mp4
	  2>失败：[AutoDownload]|Err:7|msg:1528855254  -->方法类型、错误码、文件名时间戳

http://192.168.213.133/autoDownload?channelNum=1

按时间下载接口
必填参数：channelNum<int> startTime<long> endTime<long>
返回结果：
	  1>成功：20180613-094820.mp4
	  2>失败：[DownloadByTime]|Err:7|msg:1528855254  -->方法类型、错误码、文件名时间戳

http://192.168.213.133/downloadByTime?channelNum=1&startTime=1528791412&endTime=1528791712

获取图片列表接口
http://192.168.213.133/picList/

获取视频列表接口
http://192.168.213.133/videoList/

注:如自动下载视频失败时，可通过按时间下载接口再次下载

接口说明
-----------------------------------------------------------------------------结束
*/

// Formal 全局变量
//var IPAddress string = "www.movingdt.com"
//var Port int64 = 33333
//var Account string = "admin" 
//var Password string = "a1234567"

//Test 全局变量
var IPAddress string = "192.168.88.64"
var Port int64 = 8000
var Account string = "admin" 
var Password string = "Abc123456"


/**
	自动截图在线
**/
func AutoSnap(w http.ResponseWriter, r *http.Request) {

    r.ParseForm()  //解析参数，默认是不会解析的
    /*fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
      fmt.Println("path", r.URL.Path)
      fmt.Println("scheme", r.URL.Scheme)
      fmt.Println(r.Form["channelNum"])
      for k, v := range r.Form {
          fmt.Println("key:", k)
          fmt.Println("val:", strings.Join(v, ""))
    }*/

    ip := r.Form["ip"]
    port := r.Form["port"]
    acc := r.Form["account"]
    pass := r.Form["password"]

    if ip==nil||port==nil||acc==nil||pass==nil{
    	logger.Error("Err:-990|msg:Parameter Error-VideoInfo")
		fmt.Fprintf(w, "Err:-990|msg:Parameter Error-VideoInfo")
		return
    }else{
    	IPAddress = ip[0]
    	Port, _ = strconv.ParseInt(port[0], 10, 64)
    	Account = acc[0]
    	Password = pass[0]
    }


    s := r.Form["channelNum"]
	if s==nil{
		logger.Error("Err:-997|msg:Parameter Error-channelNum")
		fmt.Fprintf(w, "Err:-997|msg:Parameter Error-channelNum")
		return
	}
	
    curSec := time.Now().Unix()
	tm := time.Unix(curSec, 0)
	filename := tm.In(cstSh).Format("20060102-150405.jpeg")
	f1 := "c"+s[0]+"-"+filename
	filename = "./img/"+f1
	fmt.Println("fileName:",filename)
	
	cstring := C.CString(filename)
	
	logger.Info("\n[========AutoSnap "+time.Unix(time.Now().Unix(), 0).In(cstSh).Format("20060102-15:04:05")+"=======")
	logger.Info("VideoInfo-ip:"+IPAddress+"-port:"+strconv.FormatInt(Port,10)+"-account:"+Account+"-password:"+Password)
	logger.Info("Parameter#channel:"+s[0]+" fileName:"+filename)
	ss := strings.Join(s,"")

	b,error := strconv.Atoi(ss)
	if error != nil{
		logger.Error("Err:-996|msg:String to int err")
		fmt.Fprintf(w, "Err:-996|msg:字符串转换成整数失败")
		return
	}		

	cint := C.int(b)
	result := C.AutoSnap(C.CString(IPAddress),C.long(Port),C.CString(Account),C.CString(Password),cstring,cint)
    fmt.Println(result)
	if result == 0 {
		logger.Info("Success|msg:"+f1)
		fmt.Fprintf(w, "Err:0|msg:"+f1)	
	}else {
		Err := strconv.Itoa(int(result)) 
		logger.Error("Err:"+Err+"|msg:"+strconv.FormatInt(curSec,10))
		fmt.Fprintf(w, "Err:"+Err+"|msg:"+strconv.FormatInt(curSec,10))	
	}
}

/**
	按时间截图 下载
**/
func AutoSnap2(w http.ResponseWriter, r *http.Request) {
    
    r.ParseForm()  //解析参数，默认是不会解析的


    ip := r.Form["ip"]
    port := r.Form["port"]
    acc := r.Form["account"]
    pass := r.Form["password"]

    if ip==nil||port==nil||acc==nil||pass==nil{
    	logger.Error("Err:-990|msg:Parameter Error-VideoInfo")
		fmt.Fprintf(w, "Err:-990|msg:Parameter Error-VideoInfo")
		return
    }else{
    	IPAddress = ip[0]
    	Port, _ = strconv.ParseInt(port[0], 10, 64)
    	Account = acc[0]
    	Password = pass[0]
    }

	//获取通道号----------------------------开始
    channel := r.Form["channelNum"]
    if channel==nil{
    	logger.Error("Err:-995|msg:Parameter Error-channelNum")
    	fmt.Fprintf(w, "Err:-995|msg:Parameter Error-channelNum")
		return
    }
    curSec := time.Now().Unix()
    logger.Info("\n[========AutoSnap2 "+time.Unix(time.Now().Unix(), 0).In(cstSh).Format("20060102-15:04:05")+"=======")
    logger.Info("VideoInfo-ip:"+IPAddress+"-port:"+strconv.FormatInt(Port,10)+"-account:"+Account+"-password:"+Password)
    logger.Info("Parameter#channel:"+channel[0]+",timestamp:"+strconv.FormatInt(curSec,10))
	channel_s := strings.Join(channel,"")

	b,error := strconv.Atoi(channel_s)
	if error != nil{
		logger.Error("Err:-997|msg:String to int err")
		fmt.Fprintf(w, "Err:-997|msg:字符串转换成整数失败")
		return
	}
	channelNum := C.int(b)
	//获取通道号----------------------------结束	

	//curSec := time.Now().Unix()
	tm := time.Unix(curSec, 0)

	//var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	//获取文件名
	filename1 := "c"+channel[0]+"-"+tm.In(cstSh).Format("20060102-150405")
	filename := "./video/avi/"+filename1+".avi"
	FileName_C := C.CString(filename)

	cur := tm.In(cstSh).Format("2006-01-02 15:04:05")
	fmt.Println(cur)
	s := strings.Split(cur, " ")
	var str1 = s[0]
	var str2 = s[1]
	ymd := strings.Split(str1, "-")  
	fmt.Print("start--->year:",ymd[0],"month:",ymd[1],"day:",ymd[2])
	hms := strings.Split(str2, ":") 
	fmt.Println("hour:",hms[0],"minute:",hms[1],"second:",hms[2])
	Year,_:= strconv.ParseInt(ymd[0], 10, 64)
	Month,_:= strconv.Atoi(ymd[1])
	Day,_:= strconv.Atoi(ymd[2])
	Hour,_:= strconv.Atoi(hms[0])
	Minute,_:= strconv.Atoi(hms[1])
	Second,_:= strconv.Atoi(hms[2])

	//格式化起始时间
	timestamp_start := curSec-1
	
	tm2 := time.Unix(timestamp_start, 0)

	startTime := tm2.In(cstSh).Format("2006-01-02 15:04:05")
	fmt.Println("startTime:",startTime)
	s2 := strings.Split(startTime, " ")
	var str1_start = s2[0]
	var str2_start = s2[1]
	ymd2 := strings.Split(str1_start, "-")  
	fmt.Print("end--->year:",ymd2[0],"month:",ymd2[1],"day:",ymd2[2])
	hms2 := strings.Split(str2_start, ":") 
	fmt.Println("hour:",hms2[0],"minute:",hms2[1],"second:",hms2[2])
	startYear,_:= strconv.ParseInt(ymd2[0], 10, 64)
	startMonth,_:= strconv.Atoi(ymd2[1])
	startDay,_:= strconv.Atoi(ymd2[2])
	startHour,_:= strconv.Atoi(hms2[0])
	startMinute,_:= strconv.Atoi(hms2[1])
	startSecond,_:= strconv.Atoi(hms2[2])
	//下载视频时需要传递起始时间和截止时间
	result :=C.DownloadByTime(C.CString(IPAddress),C.long(Port),C.CString(Account),C.CString(Password),FileName_C,channelNum,C.long(startYear),C.int(startMonth),C.int(startDay),C.int(startHour),C.int(startMinute),C.int(startSecond),C.long(Year),C.int(Month),C.int(Day),C.int(Hour),C.int(Minute),C.int(Second))
	
	if result == 0 {
		srcFileName := "/opt/go/FormalMonitorService/golangMonitor/linux64/proj/video/avi/"+filename1+".avi"
		//outputfilename := "/opt/go/FormalMonitorService/golangMonitor/linux64/proj/video/"+filename1+".mp4"
		outputfilename := "/opt/go/FormalMonitorService/golangMonitor/linux64/proj/img/"+filename1+".jpeg"

		//获取下载文件大小
		fileSize := getSize(srcFileName,filename1)
		if fileSize==0{
			startCover := time.Now()
			//删除原视频
	    	param := "rm "+srcFileName
	    	logger.Info("parm : ", param)
	    	exec.Command("bash", "-c", param).CombinedOutput()
	        //fmt.Println(" out file : ", outputfilename, " exec time  ", time.Now().Sub(startCover).Seconds())
	        logger.Info(" out file:", outputfilename, " exec time:"+FloatToString(time.Now().Sub(startCover).Seconds()))
	        logger.Error("Err:-994|msg:"+strconv.FormatInt(curSec,10)+"|errTxt:File not found!")
	        fmt.Fprintf(w, "Err:-994|msg:"+strconv.FormatInt(curSec,10)+"|errTxt:File not found!")
			return 
		}

		//param := "ffmpeg -i " + srcFileName + " " + outputfilename//视频转码
		param := "ffmpeg -i " + srcFileName + " -y -f image2 -t 0.001 -s 1280x720 "+outputfilename//视频截图
	    startCover := time.Now()
	    _, err := exec.Command("bash", "-c", param).CombinedOutput()
	    if err != nil {
		     logger.Error("Err:-994|msg:"+strconv.FormatInt(curSec,10)+"|errTxt:"+err.Error())
		     fmt.Fprintf(w, "Err:-994|msg:"+strconv.FormatInt(curSec,10)+"|errTxt:"+err.Error())
	    } else {
	    	//删除原视频
	    	param = "rm "+srcFileName
	    	logger.Info("parm : ", param)
	    	exec.Command("bash", "-c", param).CombinedOutput()
	        //fmt.Println(" out file : ", outputfilename, " exec time  ", time.Now().Sub(startCover).Seconds())
	        logger.Info(" out file:", outputfilename, " exec time:"+FloatToString(time.Now().Sub(startCover).Seconds()))
	        logger.Info("Success|msg:"+filename1+".jpeg")
	        fmt.Fprintf(w, "Err:0|msg:"+filename1+".jpeg")	
	    }
	}else {
		Err := strconv.Itoa(int(result)) 
		logger.Error("Err:"+Err+"|msg:"+strconv.FormatInt(curSec,10))
		fmt.Fprintf(w, "Err:"+Err+"|msg:"+strconv.FormatInt(curSec,10))
	}
}

/**
	按时间截图
**/
func SnapByTime(w http.ResponseWriter, r *http.Request) {
    
    r.ParseForm()  //解析参数，默认是不会解析的

    ip := r.Form["ip"]
    port := r.Form["port"]
    acc := r.Form["account"]
    pass := r.Form["password"]

    if ip==nil||port==nil||acc==nil||pass==nil{
    	logger.Error("Err:-990|msg:Parameter Error-VideoInfo")
		fmt.Fprintf(w, "Err:-990|msg:Parameter Error-VideoInfo")
		return
    }else{
    	IPAddress = ip[0]
    	Port, _ = strconv.ParseInt(port[0], 10, 64)
    	Account = acc[0]
    	Password = pass[0]
    }

	//获取通道号----------------------------开始
    channel := r.Form["channelNum"]
    curTime := r.Form["curTime"]//参数为时间戳秒数+通道号(防止同一时间多辆车请求下载文件同名)
    if channel==nil||curTime==nil{
    	logger.Error("Err:-995|msg:Parameter Error-channelNum or curTime")
    	fmt.Fprintf(w, "Err:-995|msg:Parameter Error-channelNum or curTime")
		return
    }
    logger.Info("\n[========SnapByTime "+time.Unix(time.Now().Unix(), 0).In(cstSh).Format("20060102-15:04:05")+"=======")
    logger.Info("VideoInfo-ip:"+IPAddress+"-port:"+strconv.FormatInt(Port,10)+"-account:"+Account+"-password:"+Password)
    logger.Info("Parameter#channel:"+channel[0]+",timestamp:"+curTime[0])
	channel_s := strings.Join(channel,"")

	b,error := strconv.Atoi(channel_s)
	if error != nil{
		logger.Error("Err:-997|msg:String to int err")
		fmt.Fprintf(w, "Err:-997|msg:字符串转换成整数失败")
		return
	}
	channelNum := C.int(b)
	//获取通道号----------------------------结束	


	//获取文件名---------------------------开始
	curTimeArr := curTime[0]
	timestamp, err:= strconv.ParseInt(curTimeArr, 10, 64)
	if err!=nil{
		logger.Error("Err:-997|msg:String to int err")
		fmt.Fprintf(w, "Err:-997|msg:字符串转换成整数失败")
		return
	}

	curNow := time.Now().Unix()
	if timestamp>curNow{
		logger.Error("Err:-999|msg:timestamp can`t over curTime")
		fmt.Fprintf(w, "Err:-999|msg:timestamp can`t over curTime")
		return
	}

	//格式化起始时间
	tm := time.Unix(timestamp, 0)
	//var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	//获取文件名
	filename1 := "c"+channel[0]+"-"+tm.In(cstSh).Format("20060102-150405")
	filename := "./video/avi/"+filename1+".avi"
	FileName_C := C.CString(filename)

	cur := tm.In(cstSh).Format("2006-01-02 15:04:05")
	fmt.Println(cur)
	s := strings.Split(cur, " ")
	var str1 = s[0]
	var str2 = s[1]
	ymd := strings.Split(str1, "-")  
	fmt.Print("start--->year:",ymd[0],"month:",ymd[1],"day:",ymd[2])
	hms := strings.Split(str2, ":") 
	fmt.Println("hour:",hms[0],"minute:",hms[1],"second:",hms[2])
	Year,_:= strconv.ParseInt(ymd[0], 10, 64)
	Month,_:= strconv.Atoi(ymd[1])
	Day,_:= strconv.Atoi(ymd[2])
	Hour,_:= strconv.Atoi(hms[0])
	Minute,_:= strconv.Atoi(hms[1])
	Second,_:= strconv.Atoi(hms[2])

	//格式化起始时间
	timestamp_start := timestamp-1
	
	tm2 := time.Unix(timestamp_start, 0)

	startTime := tm2.In(cstSh).Format("2006-01-02 15:04:05")
	fmt.Println("startTime:",startTime)
	s2 := strings.Split(startTime, " ")
	var str1_start = s2[0]
	var str2_start = s2[1]
	ymd2 := strings.Split(str1_start, "-")  
	fmt.Print("end--->year:",ymd2[0],"month:",ymd2[1],"day:",ymd2[2])
	hms2 := strings.Split(str2_start, ":") 
	fmt.Println("hour:",hms2[0],"minute:",hms2[1],"second:",hms2[2])
	startYear,_:= strconv.ParseInt(ymd2[0], 10, 64)
	startMonth,_:= strconv.Atoi(ymd2[1])
	startDay,_:= strconv.Atoi(ymd2[2])
	startHour,_:= strconv.Atoi(hms2[0])
	startMinute,_:= strconv.Atoi(hms2[1])
	startSecond,_:= strconv.Atoi(hms2[2])
	//下载视频时需要传递起始时间和截止时间
	result :=C.DownloadByTime(C.CString(IPAddress),C.long(Port),C.CString(Account),C.CString(Password),FileName_C,channelNum,C.long(startYear),C.int(startMonth),C.int(startDay),C.int(startHour),C.int(startMinute),C.int(startSecond),C.long(Year),C.int(Month),C.int(Day),C.int(Hour),C.int(Minute),C.int(Second))
	
	if result == 0 {
		srcFileName := "/opt/go/FormalMonitorService/golangMonitor/linux64/proj/video/avi/"+filename1+".avi"
		//outputfilename := "/opt/go/FormalMonitorService/golangMonitor/linux64/proj/video/"+filename1+".mp4"
		outputfilename := "/opt/go/FormalMonitorService/golangMonitor/linux64/proj/img/"+filename1+".jpeg"

		//获取下载文件大小
		fileSize := getSize(srcFileName,filename1)
		if fileSize==0{
			startCover := time.Now()
			//删除原视频
	    	param := "rm "+srcFileName
	    	logger.Info("parm : ", param)
	    	exec.Command("bash", "-c", param).CombinedOutput()
	        //fmt.Println(" out file : ", outputfilename, " exec time  ", time.Now().Sub(startCover).Seconds())
	        logger.Info(" out file:", outputfilename, " exec time:"+FloatToString(time.Now().Sub(startCover).Seconds()))
	        logger.Error("Err:-994|msg:"+strconv.FormatInt(timestamp,10)+"|errTxt:File not found!")
	        fmt.Fprintf(w, "Err:-994|msg:"+strconv.FormatInt(timestamp,10)+"|errTxt:File not found!")
			return 
		}

		//param := "ffmpeg -i " + srcFileName + " " + outputfilename//视频转码
		param := "ffmpeg -i " + srcFileName + " -y -f image2 -t 0.001 -s 1280x720 "+outputfilename//视频截图
	    startCover := time.Now()
	    _, err := exec.Command("bash", "-c", param).CombinedOutput()
	    if err != nil {
		     logger.Error("Err:-994|msg:"+strconv.FormatInt(timestamp,10)+"|errTxt:"+err.Error())
		     fmt.Fprintf(w, "Err:-994|msg:"+strconv.FormatInt(timestamp,10)+"|errTxt:"+err.Error())
	    } else {
	    	//删除原视频
	    	param = "rm "+srcFileName
	    	logger.Info("parm : ", param)
	    	exec.Command("bash", "-c", param).CombinedOutput()
	        //fmt.Println(" out file : ", outputfilename, " exec time  ", time.Now().Sub(startCover).Seconds())
	        logger.Info(" out file:", outputfilename, " exec time:"+FloatToString(time.Now().Sub(startCover).Seconds()))
	        logger.Info("Success|msg:"+filename1+".jpeg")
	        fmt.Fprintf(w, "Err:0|msg:"+filename1+".jpeg")	
	    }
	}else {
		Err := strconv.Itoa(int(result)) 
		logger.Error("Err:"+Err+"|msg:"+strconv.FormatInt(timestamp,10))
		fmt.Fprintf(w, "Err:"+Err+"|msg:"+strconv.FormatInt(timestamp,10))
	}
}

/*
	channelNum为必填参数 int类型
	成功-20180612-150944.mp4|视频下载成功！
	失败-错误码
*/

/**
	自动下载视频
**/
func AutoDownload(w http.ResponseWriter, r *http.Request) {
    
	tCount := 300
	t64 := int64(tCount)
    r.ParseForm()  //解析参数，默认是不会解析的


    ip := r.Form["ip"]
    port := r.Form["port"]
    acc := r.Form["account"]
    pass := r.Form["password"]

    if ip==nil||port==nil||acc==nil||pass==nil{
    	logger.Error("Err:-990|msg:Parameter Error-VideoInfo")
		fmt.Fprintf(w, "Err:-990|msg:Parameter Error-VideoInfo")
		return
    }else{
    	IPAddress = ip[0]
		Port, _ = strconv.ParseInt(port[0], 10, 64)
    	Account = acc[0]
    	Password = pass[0]
    }

	//获取通道号----------------------------开始
    channel := r.Form["channelNum"]
    timeCount := r.Form["timeCount"]//参数为时间戳秒数+通道号(防止同一时间多辆车请求下载文件同名)
    if channel==nil{
    	logger.Error("Err:-995|msg:Parameter Error_channel")
    	fmt.Fprintf(w, "Err:-995|msg:Parameter Error_channel")
		return
    }
    logger.Info("\n[========AutoDownload "+time.Unix(time.Now().Unix(), 0).In(cstSh).Format("20060102-15:04:05")+"=======")
    logger.Info("VideoInfo-ip:"+IPAddress+"-port:"+strconv.FormatInt(Port,10)+"-account:"+Account+"-password:"+Password)
    logger.Info("Parameter#channel:"+channel[0])
    if timeCount!=nil{
		ttCount,err := strconv.ParseInt(timeCount[0], 10, 64)
		if err !=nil{
			logger.Error("Err:-995|msg:Parameter Error_timeCount")
			fmt.Fprintf(w, "Err:-995|msg:Parameter Error_timeCount")
			return
		}
		t64 = ttCount
    }
	channel_s := strings.Join(channel,"")

	b,error := strconv.Atoi(channel_s)
	if error != nil{
		//fmt.Println("字符串转换成整数失败")
		logger.Error("Err:-997|msg:String to int err")
		fmt.Fprintf(w, "Err:-997|msg:字符串转换成整数失败")
		return
	}
	channelNum := C.int(b)
	//获取通道号----------------------------结束	

	//var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海

	//格式化截止时间
	curSec := time.Now().Unix()
	tm := time.Unix(curSec, 0)
	cur := tm.In(cstSh).Format("2006-01-02 15:04:05")
	fmt.Println("endTime:",cur)
	s := strings.Split(cur, " ")
	var str1 = s[0]
	var str2 = s[1]
	ymd := strings.Split(str1, "-")  
	fmt.Print("end--->year:",ymd[0]," month:",ymd[1]," day:",ymd[2])
	hms := strings.Split(str2, ":") 
	fmt.Println("  hour:",hms[0],"minute:",hms[1],"second:",hms[2])
	Year,_:= strconv.ParseInt(ymd[0], 10, 64) 
	Month,_:= strconv.Atoi(ymd[1])
	Day,_:= strconv.Atoi(ymd[2])
	Hour,_:= strconv.Atoi(hms[0])
	Minute,_:= strconv.Atoi(hms[1])
	Second,_:= strconv.Atoi(hms[2])

	//格式化起始时间
	start := curSec-t64
	tm2 := time.Unix(start, 0)

	//获取文件名
	filename := tm2.In(cstSh).Format("20060102-150405")
	filename1 := "c"+channel[0]+"-"+filename
	filename = "./video/avi/"+filename1+".avi"

	startTime := tm2.In(cstSh).Format("2006-01-02 15:04:05")
	fmt.Println("startTime:",startTime)
	s2 := strings.Split(startTime, " ")
	var str2_start = s2[1]
	hms2 := strings.Split(str2_start, ":") 
	fmt.Println("start--->hour:",hms2[0],"minute:",hms2[1],"second:",hms2[2])
	startHour,_:= strconv.Atoi(hms2[0])
	startMinute,_:= strconv.Atoi(hms2[1])
	startSecond,_:= strconv.Atoi(hms2[2])
	//下载视频时需要传递起始时间和截止时间
	result := C.AutoDownload(C.CString(IPAddress),C.long(Port),C.CString(Account),C.CString(Password),C.CString(filename),channelNum,C.long(Year),C.int(Month),C.int(Day),C.int(startHour),C.int(startMinute),C.int(startSecond),C.int(Hour),C.int(Minute),C.int(Second))
	if result == 0 {
		srcFileName := "/opt/go/FormalMonitorService/golangMonitor/linux64/proj/video/avi/"+filename1+".avi"
		outputfilename := "/opt/go/FormalMonitorService/golangMonitor/linux64/proj/video/"+filename1+".mp4"
		//获取下载文件大小
		fileSize := getSize(srcFileName,filename1)
		if fileSize==0{
			startCover := time.Now()
			//删除原视频
	    	param := "rm "+srcFileName
	    	logger.Info("parm : ", param)
	    	exec.Command("bash", "-c", param).CombinedOutput()
	        //fmt.Println(" out file : ", outputfilename, " exec time  ", time.Now().Sub(startCover).Seconds())
	        logger.Info(" out file:", outputfilename, " exec time:"+FloatToString(time.Now().Sub(startCover).Seconds()))
	        logger.Error("Err:-994|msg:"+strconv.FormatInt(curSec,10)+"|errTxt:File not found!")
			fmt.Fprintf(w, "Err:-994|msg:"+strconv.FormatInt(curSec,10)+"|errTxt:File not found!")
			return 
		}
		//go ConvertVideo(srcFileName,outputfilename)//视频转码
		logger.Info("Success|msg:"+filename1+".mp4")
	    fmt.Fprintf(w, "Err:0|msg:"+filename1+".mp4")
	}else {
		Err := strconv.Itoa(int(result)) 
		logger.Error("Err:"+Err+"|msg:"+strconv.FormatInt(start,10))
		fmt.Fprintf(w, "Err:"+Err+"|msg:"+strconv.FormatInt(start,10))
	}
}

/**
	按时间下载视频
**/
func DownloadByTime(w http.ResponseWriter, r *http.Request) {

    r.ParseForm()  //解析参数，默认是不会解析的

    ip := r.Form["ip"]
    port := r.Form["port"]
    acc := r.Form["account"]
    pass := r.Form["password"]

    if ip==nil||port==nil||acc==nil||pass==nil{
    	logger.Error("Err:-990|msg:Parameter Error-VideoInfo")
		fmt.Fprintf(w, "Err:-990|msg:Parameter Error-VideoInfo")
		return
    }else{
    	IPAddress = ip[0]
    	Port, _ = strconv.ParseInt(port[0], 10, 64)
    	Account = acc[0]
    	Password = pass[0]
    }

	//获取通道号----------------------------开始
    channel := r.Form["channelNum"]
    startTime := r.Form["startTime"]//参数为时间戳秒数+通道号(防止同一时间多辆车请求下载文件同名)
    endTime_s := r.Form["endTime"]
    if channel==nil||startTime==nil||endTime_s==nil{
    	logger.Error("Err:-996|msg:Parameter Error-channelNum_startTime_endTime")
    	fmt.Fprintf(w, "Err:-996|msg:Parameter Error-channelNum_startTime_endTime")
		return
    }
    logger.Info("\n[========DownloadByTime "+time.Unix(time.Now().Unix(), 0).In(cstSh).Format("20060102-15:04:05")+"=======")
    logger.Info("VideoInfo-ip:"+IPAddress+"-port:"+strconv.FormatInt(Port,10)+"-account:"+Account+"-password:"+Password)
    logger.Info("Parameter#channel:"+channel[0]+",startTime:"+startTime[0]+",endTime:"+endTime_s[0])
	channel_s := strings.Join(channel,"")

	b,error := strconv.Atoi(channel_s)
	if error != nil{
		//fmt.Println("字符串转换成整数失败")
		logger.Error("Err:-997|msg:String to int err")
		fmt.Fprintf(w, "Err:-997|msg:字符串转换成整数失败")
		return
	}
	channelNum := C.int(b)
	//获取通道号----------------------------结束	


	//获取文件名---------------------------开始
	//startTime := r.Form["startTime"]//参数为时间戳秒数+通道号(防止同一时间多辆车请求下载文件同名)
	startArr := startTime[0]
	timestamp, err:= strconv.ParseInt(startArr, 10, 64)
	if err!=nil{
		//fmt.Println("字符串转换成整数失败")
		logger.Error("Err:-997|msg:String to int err")
		fmt.Fprintf(w, "Err:-997|msg:字符串转换成整数失败")
		return
	}
	//格式化起始时间
	tm := time.Unix(timestamp, 0)
	//var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	//获取文件名
	filename1 := "c"+channel[0]+"-"+tm.In(cstSh).Format("20060102-150405")
	filename := "./video/avi/"+filename1+".avi"
	FileName_C := C.CString(filename)

	cur := tm.In(cstSh).Format("2006-01-02 15:04:05")
	fmt.Println(cur)
	s := strings.Split(cur, " ")
	var str1 = s[0]
	var str2 = s[1]
	ymd := strings.Split(str1, "-")  
	fmt.Print("start--->year:",ymd[0],"month:",ymd[1],"day:",ymd[2])
	hms := strings.Split(str2, ":") 
	fmt.Println("hour:",hms[0],"minute:",hms[1],"second:",hms[2])
	Year,_:= strconv.ParseInt(ymd[0], 10, 64)
	Month,_:= strconv.Atoi(ymd[1])
	Day,_:= strconv.Atoi(ymd[2])
	Hour,_:= strconv.Atoi(hms[0])
	Minute,_:= strconv.Atoi(hms[1])
	Second,_:= strconv.Atoi(hms[2])

	//格式化截止时间
	timestamp_end, err:= strconv.ParseInt(endTime_s[0], 10, 64)
	if err!=nil{
		//fmt.Println("字符串转换成整数失败")
		logger.Error("Err：-997|msg:String to int err")
		fmt.Fprintf(w, "Err：-997|msg:字符串转换成整数失败")
		return
	}
	tm2 := time.Unix(timestamp_end, 0)

	curNow := time.Now().Unix()
	if timestamp_end>curNow{
		logger.Error("Err:-999|msg:endTime can`t over curTime")
		fmt.Fprintf(w, "Err:-999|msg:endTime can`t over curTime")
		return
	}

	if timestamp>timestamp_end{
		logger.Error("Err:-999|msg:startTime can`t over endTime")
		fmt.Fprintf(w, "Err:-999|msg:startTime can`t over endTime")
		return
	}

	endTime := tm2.In(cstSh).Format("2006-01-02 15:04:05")
	fmt.Println("endTime:",endTime)
	s2 := strings.Split(endTime, " ")
	var str1_end = s2[0]
	var str2_end = s2[1]
	ymd2 := strings.Split(str1_end, "-")  
	fmt.Print("end--->year:",ymd2[0],"month:",ymd2[1],"day:",ymd2[2])
	hms2 := strings.Split(str2_end, ":") 
	fmt.Println("hour:",hms2[0],"minute:",hms2[1],"second:",hms2[2])
	endYear,_:= strconv.ParseInt(ymd2[0], 10, 64)
	endMonth,_:= strconv.Atoi(ymd2[1])
	endDay,_:= strconv.Atoi(ymd2[2])
	endHour,_:= strconv.Atoi(hms2[0])
	endMinute,_:= strconv.Atoi(hms2[1])
	endSecond,_:= strconv.Atoi(hms2[2])
	//下载视频时需要传递起始时间和截止时间
	result :=C.DownloadByTime(C.CString(IPAddress),C.long(Port),C.CString(Account),C.CString(Password),FileName_C,channelNum,C.long(Year),C.int(Month),C.int(Day),C.int(Hour),C.int(Minute),C.int(Second),C.long(endYear),C.int(endMonth),C.int(endDay),C.int(endHour),C.int(endMinute),C.int(endSecond))
	
	if result == 0 {
		srcFileName := "/opt/go/FormalMonitorService/golangMonitor/linux64/proj/video/avi/"+filename1+".avi"
		outputfilename := "/opt/go/FormalMonitorService/golangMonitor/linux64/proj/video/"+filename1+".mp4"

		//获取下载文件大小
		fileSize := getSize(srcFileName,filename1)
		if fileSize==0{
			startCover := time.Now()
			//删除原视频
	    	param := "rm "+srcFileName
	    	logger.Info("parm : ", param)
	    	exec.Command("bash", "-c", param).CombinedOutput()
	        //fmt.Println(" out file : ", outputfilename, " exec time  ", time.Now().Sub(startCover).Seconds())
	        logger.Info(" out file:", outputfilename, " exec time:"+FloatToString(time.Now().Sub(startCover).Seconds()))
	        logger.Error("Err:-994|msg:"+strconv.FormatInt(timestamp,10)+"|errTxt:File not found!")
			fmt.Fprintf(w, "Err:-994|msg:"+strconv.FormatInt(timestamp,10)+"|errTxt:File not found!")
			return 
		}

		//go ConvertVideo(srcFileName,outputfilename)//视频转码
	    logger.Info("Success|msg:"+filename1+".mp4")
	    fmt.Fprintf(w, "Err:0|msg:"+filename1+".mp4")
	}else {
		Err := strconv.Itoa(int(result)) 
		logger.Error("Err:"+Err+"|msg:"+strconv.FormatInt(timestamp,10))
		fmt.Fprintf(w, "Err:"+Err+"|msg:"+strconv.FormatInt(timestamp,10))
	}
}
/**
	获取视频格式是否已转换
**/
func IsConvert(w http.ResponseWriter, r *http.Request) {

    r.ParseForm()  //解析参数，默认是不会解析的
	
	fileName := r.Form["fileName"]
	if fileName==nil{
		logger.Error("Err:-997|msg:Parameter Error-fileName")
		fmt.Fprintf(w, "Err:-997|msg:Parameter Error-fileName")
		return
	}
	logger.Info("\n[========IsConvert "+time.Unix(time.Now().Unix(), 0).In(cstSh).Format("20060102-15:04:05")+"=======")
	logger.Info("Parameter#fileName:"+fileName[0])
	path := "/opt/go/FormalMonitorService/golangMonitor/linux64/proj/video/avi"

	fName := fileName[0]
	arr := strings.Split(fName,".")
	if len(arr)!=2{
		logger.Error("Err:-997|msg:Parameter fileName Format err")
		fmt.Fprintf(w, "Err:-997|msg:Parameter fileName Format err")
		return
	}

	fName = arr[0]+".avi"
	files, _ := ioutil.ReadDir(path)
	isConvert := true
	if len(files)==0{
		fmt.Fprintf(w,"Err:0|msg:true")
		logger.Info("IsConvert|msg:true")
		return
	}else{
		for _,file := range files{
			tfName := file.Name()
			if fName==tfName{
				isConvert = false
				break
			}
		}
		if !isConvert {
			logger.Info("IsConvert|msg:false")
			fmt.Fprintf(w,"Err:-1|msg:false")
		}else{
			logger.Info("IsConvert|msg:true")
			fmt.Fprintf(w,"Err:0|msg:true")
		}
	}

}

/**
	视频格式转换
**/

func ConvertVideo(srcFileName, outputfilename string){

		logger.Info("\n[======ConvertVideo "+time.Unix(time.Now().Unix(), 0).In(cstSh).Format("20060102-15:04:05")+"======")
		//param := "ffmpeg -y -i " + srcFileName + " -metadata:s:v rotate=0  -vf fps=15,setdar=dar=1 -s 480x480 -pix_fmt yuv420p -strict -2 -c:v h264 -b:v 500k -b:a 48k -ss 0 -t 300 -threads 2 " + outputfilename + ""
    	param := "ffmpeg -i "+srcFileName+" -y "+outputfilename//视频转码 -y覆盖输出
	    logger.Info("Convert#Parameter:"+param)
	    startCover := time.Now()
	    _, err := exec.Command("bash","-c",param).CombinedOutput()
	    if err != nil {
		     //fmt.Println("Error: " + err.Error())
		     logger.Error("Failure|Err:"+err.Error())
	    } else {
	    	//删除原视频
	    	param := "rm "+srcFileName
	    	logger.Info("Rmove#Parameter:"+param)
	    	_,err2 := exec.Command("bash", "-c", param).CombinedOutput()
	    	if err2==nil{
	    		CurConvertFile = ""
	    		logger.Info("Remove|Success")
				logger.Info("Success|out file:", outputfilename, " exec time:"+FloatToString(time.Now().Sub(startCover).Seconds()))
	    	}else{
	    		logger.Error("Remove|Err")
	    	}
	        
	    }
}


/**
	获取单个文件的大小
**/
func getSize(path,fileName string) int64 {

    fileInfo, err := os.Stat(path)
    if err != nil {
        //panic(err)
        //fmt.Println("file error")
        logger.Error("file error")
    }
    	fileSize := fileInfo.Size() //获取size
        logger.Info(fileName+".avi|size:"+strconv.FormatInt(fileSize,10) +"byte")
        return fileSize
}

func FloatToString(input_num float64) string {
    // to convert a float number to a string
    return strconv.FormatFloat(input_num, 'f', 6, 64)
}

/**
	获取系统cpu使用情况
**/
func getCPUSample() (idle, total uint64) {
    contents, err := ioutil.ReadFile("/proc/stat")
    if err != nil {
        return
    }
    lines := strings.Split(string(contents), "\n")
    for _, line := range(lines) {
        fields := strings.Fields(line)
        if fields[0] == "cpu" {
            numFields := len(fields)
            for i := 1; i < numFields; i++ {
                val, err := strconv.ParseUint(fields[i], 10, 64)
                if err != nil {
                    fmt.Println("Error: ", i, fields[i], err)
                }
                total += val // tally up all the numbers to get total ticks
                if i == 4 {  // idle is the 5th field in the cpu line
                    idle = val
                }
            }
            return
        }
    }
    return
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func WriteContent(IPAddress string,Port string,Account string,Password string){

	tempStr := IPAddress+"|"+Port+"|"+Account+"|"+Password
	logger.Info("\n[======WriteContent "+time.Unix(time.Now().Unix(), 0).In(cstSh).Format("20060102-15:04:05")+"======")
	var d1 = []byte(tempStr)
	err := ioutil.WriteFile("./info.txt", d1, 0666) //写入文件(字节数组)
    if err != nil {
        logger.Error("Err:WriteFile Failure!")
    }else{
    	logger.Info("Success:"+tempStr)
    }
}


func Test(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "test666")
	//WriteContent(IPAddress,Port,Account,Password)
}

var CurConvertFile string
var logger log.Logger
var cstSh *time.Location

func main() {

	logger = log.NewLogger("/opt/go/FormalMonitorService/golangMonitor/linux64/proj/goLog/", log.LoggerLevelInfo)
	CurConvertFile = ""
	cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	fmt.Println(reflect.TypeOf(cstSh))

	//视频转码定时任务
	ticker := time.NewTicker(time.Minute * 1)
	go func() {

	    for _ = range ticker.C {

			hour := time.Now().In(cstSh).Hour() //获取当前小时
			if hour<6||hour>22{
				//fmt.Println("Now is not work time !")
				//logger.Info("Now is not work time !")
				continue
			}

			fmt.Println("\n[======Ticker Alive Test "+time.Unix(time.Now().Unix(), 0).In(cstSh).Format("20060102-15:04:05")+"=CurConvertFile:["+CurConvertFile+"]======")
			logger.Info("\n[======Ticker Alive Test "+time.Unix(time.Now().Unix(), 0).In(cstSh).Format("20060102-15:04:05")+"=CurConvertFile:["+CurConvertFile+"]======")

	    	if CurConvertFile!=""{
	    		continue
	    	}
	    	
	        idle0, total0 := getCPUSample()
		    time.Sleep(3 * time.Second)
		    idle1, total1 := getCPUSample()
		    idleTicks := float64(idle1 - idle0)
		    totalTicks := float64(total1 - total0)
		    cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks
		    logger.Info("CPU Usage:"+FloatToString(cpuUsage)+"%|busy:"+FloatToString(totalTicks-idleTicks)+",total:"+FloatToString(totalTicks))
		    if cpuUsage>30{
		    	continue
		    }else{

		    	//自动删除超过一个月以外的日志文件
				logPath := "/opt/go/FormalMonitorService/golangMonitor/linux64/proj/goLog/"
		    	logs, _ := ioutil.ReadDir(logPath)
		    	if len(logs)>30{
		    		logName := logs[0].Name()
		    		fullLogName := logPath+logName
		    		//删除日志
			    	param := "rm "+fullLogName
			    	logger.Info("RmoveLog#Parameter:"+param)
			    	startCover := time.Now()
			    	_,err := exec.Command("bash", "-c", param).CombinedOutput()
			    	if err==nil{
			    		logger.Info("RemoveLog|Success")
						logger.Info("Success|"," exec time:"+FloatToString(time.Now().Sub(startCover).Seconds()))
			    	}else{
			    		logger.Error("RemoveLog|Err")
			    	}
		    	}

		    	//获取文件列表转码
		        path := "/opt/go/FormalMonitorService/golangMonitor/linux64/proj/video/avi/"
		        videoPath := "/opt/go/FormalMonitorService/golangMonitor/linux64/proj/video/"
		        files, _ := ioutil.ReadDir(path)

		        if len(files)>0{
		        	file := files[0]
		        	fileName := file.Name()
	                	if len(fileName)==0{
	                		logger.Error("FileName is null!")
	                		continue
	                	}
	                	fullName := path+fileName
	                	arr := strings.Split(fileName, ".")
	                    dataName := arr[0]
	                    CurConvertFile = dataName
	                    //获取文件大小
	                    size := getSize(fullName,dataName)
	                    if size==0{
	                    	//删除无效文件
						    param := "rm "+fullName
						    logger.Info("Rmove#Parameter:"+param)
						    _,err2 := exec.Command("bash", "-c", param).CombinedOutput()
						    if err2==nil{
						    	logger.Info("Remove|Success")
						    }else{
						    	logger.Error("Remove|Err")
						    }
	                    }else{//视频转码
	                    	outputFileName := videoPath+dataName+".mp4"
	                    	ConvertVideo(fullName,outputFileName)
	                    }
	        	}
		   }
		}
	}()


	//创建文件读写账号密码
	/*var filename = "./info.txt"
	var err1 error

	if checkFileIsExist(filename) { //如果文件存在
		_, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		if err1==nil{
			logger.Info("info.txt exit")
		}
	} else {
		_, err1 = os.Create(filename) //创建文件
		if err1==nil{
			logger.Info("info.txt not exit")
		}
	}*/


	http.HandleFunc("/test", Test) //设置访问的路由
	http.HandleFunc("/autoSnap", AutoSnap) //设置访问的路由
	http.HandleFunc("/autoSnap2", AutoSnap2) //设置访问的路由
	http.HandleFunc("/snapByTime", SnapByTime) //设置访问的路由
	http.Handle("/picList/", http.StripPrefix("/picList/", http.FileServer(http.Dir("/opt/go/FormalMonitorService/golangMonitor/linux64/proj/img/"))))
	http.HandleFunc("/autoDownload", AutoDownload) //设置访问的路由
	http.HandleFunc("/downloadByTime", DownloadByTime) //设置访问的路由
	http.HandleFunc("/isConvert", IsConvert) //设置访问的路由
	http.Handle("/videoList/", http.StripPrefix("/videoList/", http.FileServer(http.Dir("/opt/go/FormalMonitorService/golangMonitor/linux64/proj/video/"))))
	err := http.ListenAndServe(":8088", nil) //设置监听的端口
	if err != nil {
		logger.Error("Init===>ListenAndServe: ", err)
		fmt.Println("Init===>ListenAndServe: ", err)
	}
}

