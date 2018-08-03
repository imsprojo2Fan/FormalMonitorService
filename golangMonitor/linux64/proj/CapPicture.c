/*
* Copyright(C) 2010,Hikvision Digital Technology Co., Ltd 
* 
* File   name£ºCapPicture.cpp
* Discription£º
* Version    £º1.0
* Author     £ºpanyd
* Create Date£º2010_3_25
* Modification History£º
*/

//#include "public.h"
#include "../../include/Sdk.h"
#include <stdio.h>
#include <time.h>
#include <string.h>
/*******************************************************************
      Function:   Demo_Capture
   Description:   Capture picture.
     Parameter:   (IN)   none 
        Return:   0--success£¬-1--fail.   
**********************************************************************/
		
long AutoSnap(char *IPAddress,long Port,char *Account,char *Password,char *sPicFileName,int channelNum){

    long errCode;
    NET_DVR_Init();
    long lUserID;
    //login
    NET_DVR_DEVICEINFO_V30 struDeviceInfo;
    //lUserID = NET_DVR_Login_V30("192.168.88.8", 8000, "admin", "Abc123456", &struDeviceInfo);   
    lUserID = NET_DVR_Login_V30(IPAddress, Port, Account, Password, &struDeviceInfo);
  
    if (lUserID < 0)
    {
        printf("pyd1---Login error, %d\n", NET_DVR_GetLastError());
        errCode = NET_DVR_GetLastError();
        return errCode;
    }

    //
    NET_DVR_JPEGPARA strPicPara = {0};
    strPicPara.wPicQuality =0; //2;
    strPicPara.wPicSize = 0;//0xff; //0;
    int iRet;
    //iRet = NET_DVR_CaptureJPEGPicture(lUserID, struDeviceInfo.byStartChan, &strPicPara, "./ssss.jpeg");
    
    //printf("started chan number is: %d\n", struDeviceInfo.byStartDChan);    
    printf("started chan number is: %d\n", struDeviceInfo.byStartChan);    

    printf("the parameter is: %s\n", sPicFileName);    
 
   
    //struct tm *newtime;  
    //      char outfile[128];  
    //      time_t t1;  
    //      t1 = time(NULL);   
    //      newtime=localtime(&t1);  
    //      strftime( outfile, 128, "./img/%Y%m%d_%H%M%S.jpeg", newtime);   
   
    //printf("filename is %s \n", outfile);

    //iRet = NET_DVR_CaptureJPEGPicture(lUserID, struDeviceInfo.byStartChan, &strPicPara, outfile);    

    iRet = NET_DVR_CaptureJPEGPicture(lUserID, channelNum, &strPicPara, sPicFileName); 
    if (!iRet)
    {
        printf("pyd1---NET_DVR_CaptureJPEGPicture error, %d\n", NET_DVR_GetLastError());
        errCode = NET_DVR_GetLastError();
        return errCode;
    }

    system("ls -al /etc/passwd /etc/shadow");

    NET_DVR_Logout_V30(lUserID);
    NET_DVR_Cleanup();

    return 0;

}


long AutoDownload(char *IPAddress,long Port,char *Account,char *Password,char *sVideoFileName,int channelNum,long pYear,int pMonth,int pDay,int pHour_start,int pMinute_start,int pSecond_start,int pHour_end,int pMinute_end,int pSecond_end) {

        long errCode = 0;

        printf("the sVideoFileName is: %s\n", sVideoFileName);
        printf("CParameter--->year:%d",pYear);    
        printf("-month:%d",pMonth);
        printf("-day:%d",pDay);
        printf(" hour:%d",pHour_start);
        printf("-minute:%d",pMinute_start);
        printf("-second:%d\n",pSecond_start);
        printf("CParameter--->endHour:%d",pHour_end);
        printf("-endMinute:%d",pMinute_end);
        printf("-endSecond:%d\n",pSecond_end);
        //---------------------------------------
        NET_DVR_Init();
        //Demo_SDK_Version();
        NET_DVR_SetLogToFile(3, "./sdkLog");

        NET_DVR_SetConnectTime(2000, 1);
        NET_DVR_SetReconnect(10000, 1);
        //NET_DVR_SetReconnect(10000, true);
        //---------------------------------------
        long lUserID;
        NET_DVR_DEVICEINFO_V30 struDeviceInfo;
        //lUserID = NET_DVR_Login_V30("192.168.88.8", 8000, "admin", "Abc123456", &struDeviceInfo); 
        lUserID = NET_DVR_Login_V30(IPAddress, Port, Account, Password, &struDeviceInfo);
        if (lUserID < 0)
        {
        printf("Login error, %d\n", NET_DVR_GetLastError()); 
        errCode = NET_DVR_GetLastError();
        NET_DVR_Cleanup();
        return errCode;
        }
        NET_DVR_PLAYCOND struDownloadCond={0};
        struDownloadCond.dwChannel= channelNum;//struDeviceInfo.byStartChan; //1; 
        struDownloadCond.byStreamType= 1; 
        struDownloadCond.struStartTime.dwYear = pYear;
        struDownloadCond.struStartTime.dwMonth = pMonth;
        struDownloadCond.struStartTime.dwDay = pDay;
        struDownloadCond.struStartTime.dwHour= pHour_start;
        struDownloadCond.struStartTime.dwMinute = pMinute_start;
        struDownloadCond.struStartTime.dwSecond = pSecond_start;
        struDownloadCond.struStopTime.dwYear = pYear;
        struDownloadCond.struStopTime.dwMonth = pMonth;
        struDownloadCond.struStopTime.dwDay = pDay;
        struDownloadCond.struStopTime.dwHour = pHour_end;
        struDownloadCond.struStopTime.dwMinute = pMinute_end;
        struDownloadCond.struStopTime.dwSecond = pSecond_end;

        int hPlayback;
        hPlayback = NET_DVR_GetFileByTime_V40(lUserID,sVideoFileName,&struDownloadCond);
        if(hPlayback < 0)
        {
                printf("NET_DVR_GetFileByTime_V40 fail,last error %d\n",NET_DVR_GetLastError());
                errCode = NET_DVR_GetLastError();
                NET_DVR_Logout(lUserID);
                NET_DVR_Cleanup();
                return errCode;
        }

        printf("OK!\n");
        //---------------------------------------
        if(!NET_DVR_PlayBackControl_V40(hPlayback, NET_DVR_PLAYSTART, NULL, 0, NULL,NULL))
        {
                printf("Play back control failed [%d]\n",NET_DVR_GetLastError());
                errCode = NET_DVR_GetLastError();
                NET_DVR_Logout(lUserID);
                NET_DVR_Cleanup();
                return errCode;
        }

        printf("OK2!\n");

        int nPos = 0;
        for(nPos = 0; nPos < 100&&nPos>=0; nPos = NET_DVR_GetDownloadPos(hPlayback))
        {
                //printf("Be downloading1... %d %%\n",nPos);
                //sleep(5000); //millisecond
                //sleep(1000); //millisecond
        //return;
        }      
        if(!NET_DVR_StopGetFile(hPlayback))
        {
                printf("failed to stop get file [%d]\n",NET_DVR_GetLastError());
                errCode = NET_DVR_GetLastError();
                NET_DVR_Logout(lUserID);
                NET_DVR_Cleanup();
                return errCode;
        }
         if(nPos<0||nPos>100)
        {
                printf("download err [%d]\n",NET_DVR_GetLastError());
                errCode = NET_DVR_GetLastError();
                NET_DVR_Logout(lUserID);
                NET_DVR_Cleanup();
                return errCode;
        }
        //printf("Be downloading2... %d %%\n",nPos);
       
    NET_DVR_Logout(lUserID);
        NET_DVR_Cleanup();
        return 0;
}

long DownloadByTime(char *IPAddress,long Port,char *Account,char *Password,char *sPicFileName,int channelNum,long pYear,int pMonth,int pDay,int pHour_start,int pMinute_start,int pSecond_start,long endYear,int endMonth,int endDay,int pHour_end,int pMinute_end,int pSecond_end) {

        long errCode = 0;

        printf("the sPicFileName is: %s\n", sPicFileName);    
        //---------------------------------------
        NET_DVR_Init();
        //Demo_SDK_Version();
        NET_DVR_SetLogToFile(3, "./sdkLog");

        NET_DVR_SetConnectTime(2000, 1);
        NET_DVR_SetReconnect(10000, 1);
        //NET_DVR_SetReconnect(10000, true);
        //---------------------------------------
        long lUserID;
        NET_DVR_DEVICEINFO_V30 struDeviceInfo;
        //lUserID = NET_DVR_Login_V30("192.168.88.8", 8000, "admin", "Abc123456", &struDeviceInfo); 
        lUserID = NET_DVR_Login_V30(IPAddress, Port, Account, Password, &struDeviceInfo);
        if (lUserID < 0)
        {
        printf("Login error, %d\n", NET_DVR_GetLastError()); 
        errCode = NET_DVR_GetLastError();
        NET_DVR_Cleanup();
        return errCode;
        }
        NET_DVR_PLAYCOND struDownloadCond={0};
        struDownloadCond.dwChannel= channelNum;//struDeviceInfo.byStartChan; //1; 
        struDownloadCond.byStreamType= 1; 
        struDownloadCond.struStartTime.dwYear = pYear;
        struDownloadCond.struStartTime.dwMonth = pMonth;
        struDownloadCond.struStartTime.dwDay =pDay;
        struDownloadCond.struStartTime.dwHour= pHour_start;
        struDownloadCond.struStartTime.dwMinute =pMinute_start;
        struDownloadCond.struStartTime.dwSecond =pSecond_start;
        struDownloadCond.struStopTime.dwYear =endYear;
        struDownloadCond.struStopTime.dwMonth = endMonth;
        struDownloadCond.struStopTime.dwDay = endDay;
        struDownloadCond.struStopTime.dwHour = pHour_end;
        struDownloadCond.struStopTime.dwMinute = pMinute_end;
        struDownloadCond.struStopTime.dwSecond =pSecond_end;

        int hPlayback;
        hPlayback = NET_DVR_GetFileByTime_V40(lUserID, sPicFileName,&struDownloadCond);
        if(hPlayback < 0)
        {
                printf("NET_DVR_GetFileByTime_V40 fail,last error %d\n",NET_DVR_GetLastError());
                errCode = NET_DVR_GetLastError();
                NET_DVR_Logout(lUserID);
                NET_DVR_Cleanup();
                return errCode;
        }

        printf("OK!\n");
        //---------------------------------------
        if(!NET_DVR_PlayBackControl_V40(hPlayback, NET_DVR_PLAYSTART, NULL, 0, NULL,NULL))
        {
                printf("Play back control failed [%d]\n",NET_DVR_GetLastError());
                errCode = NET_DVR_GetLastError();
                NET_DVR_Logout(lUserID);
                NET_DVR_Cleanup();
                return errCode;
        }

        printf("OK2!\n");

        int nPos = 0;
        for(nPos = 0; nPos < 100&&nPos>=0; nPos = NET_DVR_GetDownloadPos(hPlayback))
        {
                printf("Be downloading1... %d %%\n",nPos);
                //sleep(5000); //millisecond
                //sleep(1000); //millisecond
        //return;
        }
      
        if(!NET_DVR_StopGetFile(hPlayback))
        {
                printf("failed to stop get file [%d]\n",NET_DVR_GetLastError());
                errCode = NET_DVR_GetLastError();
                NET_DVR_Logout(lUserID);
                NET_DVR_Cleanup();
                return errCode;
        }
         if(nPos<0||nPos>100)
        {
                printf("download err [%d]\n",NET_DVR_GetLastError());
                errCode = NET_DVR_GetLastError();
                NET_DVR_Logout(lUserID);
                NET_DVR_Cleanup();
                return errCode;
        }
        printf("Be downloading2... %d %%\n",nPos);
    
        NET_DVR_Logout(lUserID);
        NET_DVR_Cleanup();
        return 0;
}
