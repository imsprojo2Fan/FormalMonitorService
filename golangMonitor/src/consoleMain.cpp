 /*
* Copyright(C) 2010,Hikvision Digital Technology Co., Ltd 
* 
* File   name£ºconsoleMain.cpp
* Discription£º
* Version    £º1.0
* Author     £ºpanyadong
* Create Date£º2010_3_25
* Modification History£º
*/

#ifndef __APPLE__


#include <stdio.h>
#include <iostream>
#include "GetStream.h"
#include "public.h"
#include "ConfigParams.h"
#include "Alarm.h"
#include "CapPicture.h"
#include "playback.h"
#include "Voice.h"
#include "tool.h"
#include <time.h>
#include <stdlib.h>
#include <memory.h>
#ifdef _WIN32
#else
#include   <unistd.h> 

#endif

using namespace std;

int main()
{
    while(1)
{
    sleep(60*2); 
    //sleep(10);
    printf("screen capture is %s .\n", "beginning"); 
    NET_DVR_Init();
    Demo_SDK_Version();
    NET_DVR_SetLogToFile(3, "./sdkLog");
    char cUserChoose = 'r';
   
    //Login device
    NET_DVR_DEVICEINFO_V30 struDeviceInfo = {0};
    //LONG lUserID = NET_DVR_Login_V30("222.185.83.74", 8001, "admin", "1234567a", &struDeviceInfo);
    LONG lUserID = NET_DVR_Login_V30("www.movingdt.com", 33333, "admin", "a1234567", &struDeviceInfo);
    //LONG lUserID = NET_DVR_Login_V30("192.168.88.8", 8000, "admin", "a1234567", &struDeviceInfo);
    //LONG lUserID = NET_DVR_Login_V30("www.movingdt.com", 22222, "admin", "a1234567", &struDeviceInfo);
    if (lUserID < 0)
    {
	printf("pyd---Login error, %d\n", NET_DVR_GetLastError());
        //printf("Press any key to quit...\n");
        //cin>>cUserChoose;

	NET_DVR_Cleanup();
        //return HPR_ERROR;
        continue;
    }

	NET_DVR_JPEGPARA strPicPara = {0};
	strPicPara.wPicQuality = 0;//0;//2;
        strPicPara.wPicSize =0;//0xff; //0;
        int iRet;

	time_t currtime = time(NULL);  
	tm* p = localtime(&currtime);  
	char filename[100] = {0}; 
	sprintf(filename,"./%d%02d%02d%02d%02d%02d.jpeg",p->tm_year+1900,p->tm_mon+1,p->tm_mday,p->tm_hour,p->tm_min,p->tm_sec); 
	printf("filename is %s \n", filename);
	
        iRet = NET_DVR_CaptureJPEGPicture(lUserID, struDeviceInfo.byStartChan, &strPicPara, filename);
        //iRet = NET_DVR_CaptureJPEGPicture(lUserID, 48, &strPicPara, filename);
        printf("Capture picure's status is , %d\n", iRet);
    
/*********
    while ('q' != cUserChoose)
    {
        printf("\n");
        printf("Input 1, Test GetStream\n");
        printf("      2, Test Configure params\n");
        printf("      3, Test Alarm\n");
        printf("      4, Test Capture Picture\n");
        printf("      5, Test play back\n");
        printf("      6, Test Voice\n");
        printf("      7, Test SDK ability\n");
        printf("      8, Test tool interface\n");
        printf("      q, Quit.\n");
        printf("Input:");

        cin>>cUserChoose;
        switch (cUserChoose)
        {
        case '1':
            Demo_GetStream_V30(lUserID); //Get stream.
            break;
        case '2':
            Demo_ConfigParams(lUserID);  //Setting params.
            break;
        case '3':
            Demo_Alarm();         //Alarm & listen.
            break;
        case '4':
            Demo_Capture();
            break;
        case '5':
            Demo_PlayBack((int)lUserID);     //record & playback
            break;
        case '6':
            Demo_Voice();
            break;
        case '7':
            Demo_SDK_Ability();
            break;
		case '8':
			Demo_DVRIPByResolveSvr();
			break;
        default:
            break;
        }
    }
*/
    //logout
    NET_DVR_Logout_V30(lUserID);
    NET_DVR_Cleanup();
   ////sleep(60*5); 
}
    return 0;
}

void dowload() {
	//---------------------------------------
	// ¿¿¿
	NET_DVR_Init(); //¿¿¿¿¿¿¿¿¿¿¿ 
	NET_DVR_SetConnectTime(2000, 1);
	NET_DVR_SetReconnect(10000, true);
	//---------------------------------------
	// ¿¿¿¿
	LONG lUserID;
	NET_DVR_DEVICEINFO_V30 struDeviceInfo;
	lUserID = NET_DVR_Login_V30("192.0.0.64", 8000, "admin", "12345", &struDeviceInfo); if (lUserID < 0)
	{
	printf("Login error, %d\n", NET_DVR_GetLastError()); NET_DVR_Cleanup();
	return;
	}
	NET_DVR_PLAYCOND struDownloadCond={0};
	struDownloadCond.dwChannel=1; 
	struDownloadCond.struStartTime.dwYear = 2017;
	struDownloadCond.struStartTime.dwMonth = 11;
	struDownloadCond.struStartTime.dwDay =22;
	struDownloadCond.struStartTime.dwHour= 8;
	struDownloadCond.struStartTime.dwMinute =30;
	struDownloadCond.struStartTime.dwSecond =0;
	struDownloadCond.struStopTime.dwYear =2017;
	struDownloadCond.struStopTime.dwMonth = 11;
	struDownloadCond.struStopTime.dwDay = 22;
	struDownloadCond.struStopTime.dwHour = 8;
	struDownloadCond.struStopTime.dwMinute = 32;
	struDownloadCond.struStopTime.dwSecond =0;

	int hPlayback;
	hPlayback = NET_DVR_GetFileByTime_V40(lUserID, "./test.mp4",&struDownloadCond);
	if(hPlayback < 0)
	{
		printf("NET_DVR_GetFileByTime_V40 fail,last error %d\n",NET_DVR_GetLastError());
		NET_DVR_Logout(lUserID);
		NET_DVR_Cleanup();
		return;
	}

	//---------------------------------------
	if(!NET_DVR_PlayBackControl_V40(hPlayback, NET_DVR_PLAYSTART, NULL, 0, NULL,NULL))
	{
		printf("Play back control failed [%d]\n",NET_DVR_GetLastError());
		NET_DVR_Logout(lUserID);
		NET_DVR_Cleanup();
		return;
	}

	int nPos = 0;
	for(nPos = 0; nPos < 100&&nPos>=0; nPos = NET_DVR_GetDownloadPos(hPlayback))
	{
		printf("Be downloading... %d %%\n",nPos);
		sleep(5000); //millisecond
	}
	if(!NET_DVR_StopGetFile(hPlayback))
	{
		printf("failed to stop get file [%d]\n",NET_DVR_GetLastError()); 
		NET_DVR_Logout(lUserID);
		NET_DVR_Cleanup();
		return;
	}
	if(nPos<0||nPos>100)
	{
		printf("download err [%d]\n",NET_DVR_GetLastError());
		NET_DVR_Logout(lUserID);
		NET_DVR_Cleanup();
		return;
	}
	printf("Be downloading... %d %%\n",nPos);
	NET_DVR_Logout(lUserID);
	NET_DVR_Cleanup();
	return;
}

#endif
