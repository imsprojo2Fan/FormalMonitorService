package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"time"
	"strconv"
)

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -L ../lib -Wl,-rpath=./:./HCNetSDKCom:../lib   -lCapPicture -lhcnetsdk
#include "../../include/Sdk.h"
#include "CapPicture.h"
*/
import "C"


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
	
    var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	filename := time.Now().In(cstSh).Format("20060102-150405.jpeg")
	filename = "./img/"+filename
	fmt.Println("fileName:",filename)
	fmt.Println("curSec:",time.Now().Unix())
	cstring := C.CString(filename)
	s := r.Form["channelNum"]
	if s==nil{
		fmt.Fprintf(w, "[AutoSnap]Err:Parameter Error-channelNum")
		return
	}
	ss := strings.Join(s,"")

	b,error := strconv.Atoi(ss)
	if error != nil{
		fmt.Println("字符串转换成整数失败")
		fmt.Fprintf(w, "[AutoSnap]Err:字符串转换成整数失败")
		return
	}		

	cint := C.int(b)
	result :=C.AutoSnap(cstring,cint)
    fmt.Println(result)
	if result == 0 {
		fmt.Fprintf(w, filename[6:26])	
	}else {
		Err := strconv.Itoa(int(result)) 
		fmt.Fprintf(w, "[AutoSnap]Err:"+Err)	
	}
}

func AutoDownload(w http.ResponseWriter, r *http.Request) {
    
    r.ParseForm()  //解析参数，默认是不会解析的
	//获取通道号----------------------------开始
    channel := r.Form["channelNum"]
    fileName := r.Form["endTime"]//参数为时间戳秒数+通道号(防止同一时间多辆车请求下载文件同名)
    if channel==nil|| fileName==nil{
    	fmt.Fprintf(w, "[AutoDownload]Err:Parameter Error")
		return
    }
	channel_s := strings.Join(channel,"")

	b,error := strconv.Atoi(channel_s)
	if error != nil{
		fmt.Println("字符串转换成整数失败")
		fmt.Fprintf(w, "[AutoDownload]Err:字符串转换成整数失败")
		return
	}
	channelNum := C.int(b)
	//获取通道号----------------------------结束	

	//filename := time.Now().Format("./video/20060102-150405.mp4")
	//FileName := C.CString(filename)

	//获取文件名---------------------------开始
	//fileName := r.Form["filaName"]//参数为时间戳秒数+通道号(防止同一时间多辆车请求下载文件同名)
	fileNameArr := fileName[0]
	FileName := "./video/"+fileNameArr+".mp4"
	fmt.Println("[AutoDownload]fileName:"+FileName)
	FileName_C := C.CString(FileName)


	timestamp, err:= strconv.ParseInt(fileNameArr, 10, 64)
	if err!=nil{
		fmt.Println("字符串转换成整数失败")
		fmt.Fprintf(w, "[AutoDownload]Err:字符串转换成整数失败")
		return
	}
	//格式化截止时间
	tm := time.Unix(timestamp, 0)
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
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
	start := timestamp-300
	tm2 := time.Unix(start, 0)
	startTime := tm2.In(cstSh).Format("2006-01-02 15:04:05")
	fmt.Println("startTime:",startTime)
	s2 := strings.Split(startTime, " ")
	var str2_start = s2[1]
	hms2 := strings.Split(str2_start, ":") 
	fmt.Println("start--->hour:",hms2[0],"minute:",hms2[1],"second:",hms2[2])
	endHour,_:= strconv.Atoi(hms2[0])
	endMinute,_:= strconv.Atoi(hms2[1])
	endSecond,_:= strconv.Atoi(hms2[2])
	//下载视频时需要传递起始时间和截止时间
	result :=C.AutoDownload(FileName_C,channelNum,C.long(Year),C.int(Month),C.int(Day),C.int(Hour),C.int(Minute),C.int(Second),C.int(endHour),C.int(endMinute),C.int(endSecond))
	if result == 0 {
		fmt.Fprintf(w,"视频下载成功！")	
	}else {
		Err := strconv.Itoa(int(result)) 
		fmt.Fprintf(w, "[AutoDownload]Err:"+Err)
	}
}

/*func AutoDownload2(w http.ResponseWriter, r *http.Request) {
    
	result := C.AutoDownload2()
	if result == 0 {
		fmt.Fprintf(w,"视频下载成功！")	
	}else {
		Err := strconv.Itoa(int(result)) 
		fmt.Fprintf(w, "[AutoDownload2]Err:"+Err)
	}
}*/

func DownloadByTime(w http.ResponseWriter, r *http.Request) {
    
    r.ParseForm()  //解析参数，默认是不会解析的
	//获取通道号----------------------------开始
    channel := r.Form["channelNum"]
    startTime := r.Form["startTime"]//参数为时间戳秒数+通道号(防止同一时间多辆车请求下载文件同名)
    endTime_s := r.Form["endTime"]
    if channel==nil||startTime==nil||endTime_s==nil{
    	fmt.Fprintf(w, "[DownloadByTime]Err:Parameter Error-channelNum|startTime|endTime")
		return
    }
	channel_s := strings.Join(channel,"")

	b,error := strconv.Atoi(channel_s)
	if error != nil{
		fmt.Println("字符串转换成整数失败")
		fmt.Fprintf(w, "[DownloadByTime]Err:字符串转换成整数失败")
		return
	}
	channelNum := C.int(b)
	//获取通道号----------------------------结束	

	//filename := time.Now().Format("./video/20060102-150405.mp4")
	//FileName := C.CString(filename)

	//获取文件名---------------------------开始
	//startTime := r.Form["startTime"]//参数为时间戳秒数+通道号(防止同一时间多辆车请求下载文件同名)
	startArr := startTime[0]
	FileName := "./video/"+startArr+".mp4"
	fmt.Println("[DownloadByTime]fileName:"+FileName)
	FileName_C := C.CString(FileName)


	timestamp, err:= strconv.ParseInt(startArr, 10, 64)
	if err!=nil{
		fmt.Println("字符串转换成整数失败")
		fmt.Fprintf(w, "[DownloadByTime]Err:字符串转换成整数失败")
		return
	}
	//格式化起始时间
	tm := time.Unix(timestamp, 0)
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
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
		fmt.Println("字符串转换成整数失败")
		fmt.Fprintf(w, "[DownloadByTime]Err:字符串转换成整数失败")
		return
	}
	tm2 := time.Unix(timestamp_end, 0)
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
	result :=C.DownloadByTime(FileName_C,channelNum,C.long(Year),C.int(Month),C.int(Day),C.int(Hour),C.int(Minute),C.int(Second),C.long(endYear),C.int(endMonth),C.int(endDay),C.int(endHour),C.int(endMinute),C.int(endSecond))
	
	if result == 0 {
		fmt.Fprintf(w,"视频下载成功！")	
	}else {
		Err := strconv.Itoa(int(result)) 
		fmt.Fprintf(w, "[DownloadByTime]Err:"+Err)
	}
}

func Test(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "test666")
}

func main() {
	http.HandleFunc("/test", Test) //设置访问的路由
	http.HandleFunc("/autoSnap", AutoSnap) //设置访问的路由
	http.Handle("/picList/", http.StripPrefix("/picList/", http.FileServer(http.Dir("/opt/go/FormalMonitorService/golangMonitor/linux64/proj/img/"))))
	http.HandleFunc("/autoDownload", AutoDownload) //设置访问的路由
	//http.HandleFunc("/autoDownloadTest", AutoDownload2) //设置访问的路由
	http.HandleFunc("/downloadByTime", DownloadByTime) //设置访问的路由
	http.Handle("/videoList/", http.StripPrefix("/videoList/", http.FileServer(http.Dir("/opt/go/FormalMonitorService/golangMonitor/linux64/proj/video/"))))
	err := http.ListenAndServe(":8088", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

