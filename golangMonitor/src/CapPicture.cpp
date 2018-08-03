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
#include "../include/Sdk.h"
#include "CapPicture.h"
#include <stdio.h>

/*******************************************************************
      Function:   Demo_Capture
   Description:   Capture picture.
     Parameter:   (IN)   none 
        Return:   0--success£¬-1--fail.   
**********************************************************************/
int Demo_Capture()
{
    NET_DVR_Init();
    long lUserID;
    //login
    NET_DVR_DEVICEINFO_V30 struDeviceInfo;
    lUserID = NET_DVR_Login_V30("222.185.83.74", 8001, "admin", "1234567a", &struDeviceInfo);
    if (lUserID < 0)
    {
        printf("pyd1---Login error, %d\n", NET_DVR_GetLastError());
        return HPR_ERROR;
    }

    //
    NET_DVR_JPEGPARA strPicPara = {0};
    strPicPara.wPicQuality =0; //2;
    strPicPara.wPicSize =10; //0;
    int iRet;
    //iRet = NET_DVR_CaptureJPEGPicture(lUserID, struDeviceInfo.byStartChan, &strPicPara, "./ssss.jpeg");
    
    printf("started chan number is: %d\n", struDeviceInfo.byStartDChan);    

    iRet = NET_DVR_CaptureJPEGPicture(lUserID, 36, &strPicPara, "./ssssd.jpeg");
    if (!iRet)
    {
        printf("pyd1---NET_DVR_CaptureJPEGPicture error, %d\n", NET_DVR_GetLastError());
        return HPR_ERROR;
    }

    //logout
    NET_DVR_Logout_V30(lUserID);
    NET_DVR_Cleanup();

    return HPR_OK;

}
