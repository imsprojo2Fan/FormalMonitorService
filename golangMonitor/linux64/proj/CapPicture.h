/*
* Copyright(C) 2010,Hikvision Digital Technology Co., Ltd 
* 
* File   name��CapPicture.h
* Discription��
* Version    ��1.0
* Author     ��panyd
* Create Date��2010_3_25
* Modification History��
*/

int AutoSnap(char *IPAddress,long Port,char *Account,char *Password,char *sPicFileName, int channelNum);

long AutoDownload(char *IPAddress,long Port,char *Account,char *Password,char *sPicFileName,int channelNum,long pYear,int pMonth,int pDay,int pHour_start,int pMinute_start,int pSecond_start,int pHour_end,int pMinute_end,int pSecond_end);

long DownloadByTime(char *IPAddress,long Port,char *Account,char *Password,char *sPicFileName,int channelNum,long pYear,int pMonth,int pDay,int pHour_start,int pMinute_start,int pSecond_start,long endYear,int endMonth,int endDay,int pHour_end,int pMinute_end,int pSecond_end);
