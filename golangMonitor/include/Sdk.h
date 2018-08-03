//#ifndef _PUBLIC_H_
//#define _PUBLIC_H_

#define  HPR_OK 0
#define  HPR_ERROR -1
#define __stdcall
#define STREAM_ID_LEN   32
 
//#ifndef _HC_NET_SDK_H_
//#define _HC_NET_SDK_H_
 
#ifdef __cplusplus
    #define NET_DVR_API extern "C"
 #endif


#define SERIALNO_LEN		    48      //序列号长度
#define NET_DVR_PLAYSTART        1//开始播放

#ifdef __cplusplus
NET_DVR_API int __stdcall NET_DVR_Init();
NET_DVR_API int __stdcall NET_DVR_Cleanup();
#endif

int __stdcall NET_DVR_Init();
int __stdcall NET_DVR_Cleanup();

typedef struct
{
    unsigned char sSerialNumber[SERIALNO_LEN];  //序列号
    unsigned char byAlarmInPortNum;                //报警输入个数
    unsigned char byAlarmOutPortNum;                //报警输出个数
    unsigned char byDiskNum;                    //硬盘个数
    unsigned char byDVRType;                    //设备类型, 1:DVR 2:ATM DVR 3:DVS ......
    unsigned char byChanNum;                    //模拟通道个数
    unsigned char byStartChan;                    //起始通道号,例如DVS-1,DVR - 1
    unsigned char byAudioChanNum;                //语音通道数
    unsigned char byIPChanNum;                    //最大数字通道个数，低位  
    unsigned char byZeroChanNum;            //零通道编码个数 //2010-01-16
    unsigned char byMainProto;            //主码流传输协议类型 0-private, 1-rtsp,2-同时支持private和rtsp
    unsigned char bySubProto;                //子码流传输协议类型0-private, 1-rtsp,2-同时支持private和rtsp
    unsigned char bySupport;        //能力，位与结果为0表示不支持，1表示支持，
    unsigned char bySupport1;        // 能力集扩充，位与结果为0表示不支持，1表示支持
    unsigned char bySupport2; /*能力，位与结果为0表示不支持，非0表示支持                            
                     bySupport2 & 0x1, 表示解码器是否支持通过URL取流解码
                     bySupport2 & 0x2,  表示支持FTPV40
                     bySupport2 & 0x4,  表示支持ANR
                     bySupport2 & 0x8,  表示支持CCD的通道参数配置
                     bySupport2 & 0x10,  表示支持布防报警回传信息（仅支持抓拍机报警 新老报警结构）
                     bySupport2 & 0x20,  表示是否支持单独获取设备状态子项
    bySupport2 & 0x40,  表示是否是码流加密设备*/
    unsigned short wDevType;              //设备型号
    unsigned char bySupport3; //能力集扩展，位与结果为0表示不支持，1表示支持
    unsigned char byMultiStreamProto;//是否支持多码流,按位表示,0-不支持,1-支持,bit1-码流3,bit2-码流4,bit7-主码流，bit-8子码流
    unsigned char byStartDChan;        //起始数字通道号,0表示无效
    unsigned char byStartDTalkChan;    //起始数字对讲通道号，区别于模拟对讲通道号，0表示无效
    unsigned char byHighDChanNum;        //数字通道个数，高位
    unsigned char bySupport4;        //能力集扩展，位与结果为0表示不支持，1表示支持
    unsigned char byLanguageType;// 支持语种能力,按位表示,每一位0-不支持,1-支持  
    unsigned char byVoiceInChanNum;   //音频输入通道数 
    unsigned char byStartVoiceInChanNo; //音频输入起始通道号 0表示无效
    unsigned char  bySupport5;  //按位表示,0-不支持,1-支持,bit0-支持多码流
    unsigned char  bySupport6;   //能力，按位表示，0-不支持,1-支持
    unsigned char  byMirrorChanNum;    //镜像通道个数，<录播主机中用于表示导播通道>
    unsigned short wStartMirrorChanNo;  //起始镜像通道号
    unsigned char bySupport7;   //能力,按位表示,0-不支持,1-支持
    unsigned char  byRes2;        //保留
}NET_DVR_DEVICEINFO_V30, *LPNET_DVR_DEVICEINFO_V30;

//图片质量
typedef struct 
{
    unsigned short    wPicSize;            
    unsigned short    wPicQuality;            /* 图片质量系数 0-最好 1-较好 2-一般 */
}NET_DVR_JPEGPARA, *LPNET_DVR_JPEGPARA;

//校时结构参数
typedef struct
{
    unsigned int dwYear;        //年
    unsigned int dwMonth;        //月
    unsigned int dwDay;        //日
    unsigned int dwHour;        //时
    unsigned int dwMinute;        //分
    unsigned int dwSecond;        //秒
}NET_DVR_TIME, *LPNET_DVR_TIME;

typedef struct tagNET_DVR_PLAYCOND
{
    unsigned int             dwChannel;
    NET_DVR_TIME     struStartTime;
    NET_DVR_TIME     struStopTime;
    unsigned char             byDrawFrame;  //0:不抽帧，1：抽帧
    unsigned char             byStreamType ; //码流类型，0-主码流 1-子码流 2-码流三
    unsigned char             byStreamID[STREAM_ID_LEN];
    unsigned char             byRes[30];    //保留
}NET_DVR_PLAYCOND, *LPNET_DVR_PLAYCOND;


//NET_DVR_API int __stdcall NET_DVR_CaptureJPEGPicture(int lUserID, int lChannel, LPNET_DVR_JPEGPARA lpJpegPara, char *sPicFileName);
#ifdef __cplusplus
NET_DVR_API int __stdcall NET_DVR_Login_V30(char *sDVRIP, unsigned short wDVRPort, char *sUserName, char *sPassword, LPNET_DVR_DEVICEINFO_V30 lpDeviceInfo);

NET_DVR_API int __stdcall NET_DVR_Logout_V30(int lUserID);

NET_DVR_API int __stdcall NET_DVR_Cleanup();

NET_DVR_API int __stdcall NET_DVR_CaptureJPEGPicture(int lUserID, int lChannel, LPNET_DVR_JPEGPARA lpJpegPara, char *sPicFileName);

NET_DVR_API unsigned int __stdcall NET_DVR_GetLastError();

NET_DVR_API int __stdcall NET_DVR_GetFileByTime_V40(int lUserID, char *sSavedFileName, LPNET_DVR_PLAYCOND  pDownloadCond);
#else

 int __stdcall NET_DVR_Login_V30(char *sDVRIP, unsigned short wDVRPort, char *sUserName, char *sPassword, LPNET_DVR_DEVICEINFO_V30 lpDeviceInfo);

 int __stdcall NET_DVR_Logout_V30(int lUserID);

 int __stdcall NET_DVR_Cleanup();

 int __stdcall NET_DVR_CaptureJPEGPicture(int lUserID, int lChannel, LPNET_DVR_JPEGPARA lpJpegPara, char *sPicFileName);

 unsigned int __stdcall NET_DVR_GetLastError();

 int __stdcall NET_DVR_GetFileByTime_V40(int lUserID, char *sSavedFileName, LPNET_DVR_PLAYCOND  pDownloadCond);
#endif
