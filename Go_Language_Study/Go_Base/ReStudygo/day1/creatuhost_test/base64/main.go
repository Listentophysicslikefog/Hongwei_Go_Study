 package main
 
 import (
     "encoding/base64"
     "fmt"
     "log"
 )

 func main() {
     input := []byte("Password1")

    // 演示base64编码
     encodeString := base64.StdEncoding.EncodeToString(input)
     fmt.Println(encodeString)
 
     // 对上面的编码结果进行base64解码
     decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
     if err != nil {
         log.Fatalln(err)
     }
     fmt.Println(string(decodeBytes))
 
     fmt.Println("///////////////////////////////")
 
     // 如果要用在url中，需要使用URLEncoding
     uEnc := base64.URLEncoding.EncodeToString([]byte(input))
     fmt.Println(uEnc)
 
    uDec, err := base64.URLEncoding.DecodeString(uEnc)
     if err != nil {
         log.Fatalln(err)
    }
     fmt.Println(string(uDec))
 }


 func FioIopsTest(ctx *InnerResourceInitContext, fioParams FioTestParams, alluhostlist []ApplyUHostInfo, iplist []string, AllUhostTestDiskNum uint32, iplistUhostAddTestDiskNum uint32, OpTrack *OperationTrack) (rc *ResponseCode) {

	ctx.INFOF("获取的uhost详细信息是$$$$$$$$$$$$$$$$ : %v", alluhostlist)

	fio_bs = " -bs=" + fioParams.FioBs
	fio_iodepth = " -iodepth=" + fioParams.FioIodepth
	fio_size = " -size=" + fioParams.IoSize
	fio_runtime = " -runtime=" + fioParams.FioRuntime
	fio_numjobs = " -numjobs=" + fioParams.FioNumjobs
	fio_filename = " -filename=/dev/"
	fio_rwmixread = " -rwmixread=" + fioParams.FioIopsRWMixWrite

	var stdoutfile string
	var stderrorfile string

	var all_test_disk_num uint32

	if AllUhostTestDiskNum == FIO_TEST_UDISK_NUMBER_ERROR && iplistUhostAddTestDiskNum == FIO_TEST_UDISK_NUMBER_ERROR {

		ctx.ERRORF("You Fio Test Udisk Number Is Error, Please Check You Config |   AllUhostTestDiskNum : %v  iplistUhostAddTestDiskNum  :  %v ", AllUhostTestDiskNum, iplistUhostAddTestDiskNum)
		rc = NewResponseCode(FIO_TEST_UDISK_NUM_IN_CONFIG_ERROR)
		rc.SetAction(fmt.Sprintf("FioBWTest"))
		return rc

	}

	all_test_disk_num = uint32(len(alluhostlist))*AllUhostTestDiskNum + uint32(len(iplist))*iplistUhostAddTestDiskNum

	var diskNum string
	var getUhostMountUdiskNum map[string]uint32
	getUhostMountUdiskNum = make(map[string]uint32, 0)

	OpTrack.CurrOpTk.IsRemoteOperation = true
	OpTrack.CurrOpTk.CurrentOperationId = 0

	for _, hostinfo := range alluhostlist {
		OpTrack.CurrOpTk.CurrentOperationDescrition = "获取虚机挂载udiak的个数"
		OpTrack.CurrOpTk.RemoteOp.Ip = hostinfo.Internal_Ip
		OpTrack.CurrOpTk.RemoteOp.Cmd = GET_MOUNT_UDISK_NUMBER
		OpTrack.CurrOpTk.RemoteOp.User = hostinfo.User
		OpTrack.CurrOpTk.RemoteOp.Password = hostinfo.Password

		diskNum, rc = RemoteCmdRequestWithResultWithRN(
			OpTrack.CurrOpTk.RemoteOp.User,
			OpTrack.CurrOpTk.RemoteOp.Password,
			OpTrack.CurrOpTk.RemoteOp.Ip,
			OpTrack.CurrOpTk.RemoteOp.Cmd)
		if rc != nil {
			ctx.ERRORF("Get Uhost Mount Udisk Number  Fail In Ip : [%v] | err: %v", hostinfo.Internal_Ip, rc)
			OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
			return rc
		}
		//去掉换行符，不然会报错，diskNum 为 ： "0\n"   ,strconv.Atoi(diskNum) 为0
		diskNum = strings.Replace(diskNum, "\n", "", -1)
		diskNumber, err := strconv.Atoi(diskNum)
		if err != nil {
			fmt.Printf("666")
		}
		OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestWithResultWithRN succeed")
		getUhostMountUdiskNum[hostinfo.Internal_Ip] = uint32(diskNumber)
		ctx.INFOF("挂载磁盘的数量为 ：%v   uhost mount udisk number  | disk number %v", getUhostMountUdiskNum[hostinfo.Internal_Ip], diskNumber)
	}

	var allTestUdiskUhostList map[string]ApplyUHostInfo
	allTestUdiskUhostList = make(map[string]ApplyUHostInfo, 0)

	//注意 AllUhostTestDiskNum 可以进入这个函数一定不为空并且有虚机信息    iplistUhostAddTestDiskNumd对应的列表可能为nil

	for _, uhostinfo := range alluhostlist {
		allTestUdiskUhostList[uhostinfo.Internal_Ip] = uhostinfo
	}

	var addTestUdiskUhostList map[string]ApplyUHostInfo
	addTestUdiskUhostList = make(map[string]ApplyUHostInfo, 0)
	if len(iplist) > 0 && iplist[0] != "" {
		for _, ip := range iplist {
			value, ok := allTestUdiskUhostList[ip]
			if ok != true {
				ctx.ERRORF("addTestUdiskUhostList  |  ip : %v | not exit in allTestUdiskUhostList, please check you config or chaos mongo ", ip)
				rc = NewResponseCode(FIO_TEST_IP_NOT_EXIT)
				rc.SetAction(fmt.Sprintf("FioBWTest"))
				return rc
			}
			addTestUdiskUhostList[ip] = value
        ctx.INFOF("Fio Test Ip addTestUdiskUhostList Exit In allTestUdiskUhostList | You addTestUdiskUhostList One Ip Is : %v", ip)
		}
	}

	var fio_cmd string
	//开始进行对所有的要测试fio的虚机的磁盘进行压测
	if AllUhostTestDiskNum != FIO_TEST_UDISK_NUMBER_ERROR {
		for _, uhostinfo := range alluhostlist {
			if getUhostMountUdiskNum[uhostinfo.Internal_Ip] < AllUhostTestDiskNum {
				//报错
				ctx.ERRORF("Fio test udisk number in uhost not enough maybe you config fio test udisk number too many | UhostMountUdiskNum : %v    AllUhostTestDiskNum  : %v", getUhostMountUdiskNum[uhostinfo.Internal_Ip], AllUhostTestDiskNum)
				rc = NewResponseCode(FIO_TEST_UDISK_NUM_NOT_ENOUGH)
				rc.SetAction(fmt.Sprintf("FioIopsTest"))
				return rc
			}
			// 开始压测
			ctx.INFOF("Start Fio Test | Now  Test Ip Is : %v  Test Mode : %v ",uhostinfo.Internal_Ip, fioParams.FioRWMode)
			var testnum uint32
			for testnum = 0; testnum < AllUhostTestDiskNum; testnum++ {
				disk_nvme := uhostinfo.UDisk_App_Info[testnum].Disk_Lable

				stdoutfile = "Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme + "_result.txt"
				stderrorfile = "Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme + "_error.txt"
				fio_threadname = "  -name=Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme

				if fioParams.FioRWMode == "randread" {

					fio_cmd = IOPS_RANDREAD_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else if fioParams.FioRWMode == "randwrite" {

					fio_cmd = IOPS_RANDWRITE_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else if fioParams.FioRWMode == "randrw" {

					fio_cmd = IOPS_RWMIX_MODE_TEST_CMD + fio_rwmixread + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else {
					ctx.ERRORF("Fio Iops test No Such Mode, Please Check You Config Params [FioRWMode] | This Mode :[%v]  Is Error ", fioParams.FioRWMode)
					rc = NewResponseCode(FIO_TEST_RWMODE_ERROR)
				    rc.SetAction(fmt.Sprintf("FioIopsTest"))
					return rc 
				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "开始对所有虚机进行fio压测"

				OpTrack.CurrOpTk.RemoteOp.User = uhostinfo.User
				OpTrack.CurrOpTk.RemoteOp.Password = uhostinfo.Password
				OpTrack.CurrOpTk.RemoteOp.Ip = uhostinfo.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = fio_cmd
				rc = RemoteCmdRequest(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Fio test start fail, please check you commond or remot uhost  | err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestNoResult failed")
					return rc
				}
				OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestNoResult succeed")

				ctx.INFOF("Uhost : [%v]  Start Fio IopsTest  Successful  This Uhost Test Udisk Nvme Is :%v",uhostinfo.Internal_Ip, disk_nvme)
			}

			ctx.INFOF("Uhost : [%v]  Start Fio IopsTest  Successful  This Uhost Test Udisk Number Is :%v",uhostinfo.Internal_Ip, testnum)

		}
		ctx.INFOF(" All Uhost Start Fio IopsTest Successful | alluhostlist : [%v]   Per Uhost  Udisk Number : %v | Will Start iplistUhostAddTestDisk Fio Test",alluhostlist, uint32(len(alluhostlist))*AllUhostTestDiskNum)
	}
	//对需要额外加盘测试的虚机的磁盘进行fio压测    仔细考虑流程和各种情况
	if iplistUhostAddTestDiskNum != FIO_TEST_UDISK_NUMBER_ERROR && len(addTestUdiskUhostList) > 0 {
		for _, value := range addTestUdiskUhostList {

			if getUhostMountUdiskNum[value.Internal_Ip] < (iplistUhostAddTestDiskNum + AllUhostTestDiskNum) {
				//报错
				ctx.ERRORF("Fio test udisk number in uhost not enough , add not enough  maybe you config fio test udisk number too many | UhostMountUdiskNum : %v    AllUhostTestDiskNum  : %v     iplistUhostAddTestDiskNum  : %v", getUhostMountUdiskNum[value.Internal_Ip], AllUhostTestDiskNum, iplistUhostAddTestDiskNum)
				rc = NewResponseCode(FIO_TEST_UDISK_NUM_NOT_ENOUGH)
				rc.SetAction(fmt.Sprintf("FioBWTest"))
				return rc
			}
			var testnum uint32
			for testnum = AllUhostTestDiskNum; testnum < (iplistUhostAddTestDiskNum + AllUhostTestDiskNum); testnum++ {
				disk_nvme := value.UDisk_App_Info[testnum].Disk_Lable
				//2>>stderr.txt 1>stdout3.txt
				stdoutfile = "Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme + "_result.txt"
				stderrorfile = "Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme + "_error.txt"
				fio_threadname = "  -name=Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme

				if fioParams.FioRWMode == "randread" {

					fio_cmd = IOPS_RANDREAD_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else if fioParams.FioRWMode == "randwrite" {

					fio_cmd = IOPS_RANDWRITE_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else if fioParams.FioRWMode == "randrw" {

					fio_cmd = IOPS_RWMIX_MODE_TEST_CMD + fio_rwmixread + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else {
					ctx.ERRORF("Fio Iops test No Such Mode, Please Check You Config Params [FioRWMode] | This Mode :[%v]  Is Error ", fioParams.FioRWMode)
					rc = NewResponseCode(FIO_TEST_RWMODE_ERROR)
				    rc.SetAction(fmt.Sprintf("FioIopsTest"))
					return rc 
				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "开始对虚机进行fio压测"

				OpTrack.CurrOpTk.RemoteOp.User = value.User
				OpTrack.CurrOpTk.RemoteOp.Password = value.Password
				OpTrack.CurrOpTk.RemoteOp.Ip = value.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = fio_cmd
				rc = RemoteCmdRequest(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Fio test start  fail, please check you commond or remot uhost  | err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestNoResult failed")
					return rc
				}
				OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestNoResult succeed")
				ctx.INFOF("Uhost : [%v]  Start Fio IopsTest  Successful  This Uhost Test Udisk Nvme Is :%v",value.Internal_Ip, disk_nvme)
			}
			ctx.INFOF("Uhost : [%v]  Start Fio IopsTest  Successful | This Uhost Test Udisk Number Is :%v",value.Internal_Ip, testnum)
		}
		ctx.INFOF(" All iplistUhostAddTestDisk Start Fio IopsTest Successful iplistUhostAddTestDisk : [%v]  |  Per Uhost  Udisk Number : %v",iplist, uint32(len(iplist))*iplistUhostAddTestDiskNum)

	}

	ctx.INFOF("AllUhostTestDisk And iplistUhostAddTestDisk Start Fio Iops Test Successful | All  Fio Test Udisk Number  is : %v ", all_test_disk_num)
	return nil

}

func GetFioIopsTestResultByTestType(ctx *InnerResourceInitContext, disktype string, fioParams FioTestParams, fiocreatdisk CreateUdiskParams, operation string, alluhostlist []ApplyUHostInfo, iplist []string, AllUhostTestDiskNum uint32, iplistUhostAddTestDiskNum uint32, OpTrack *OperationTrack) (rc *ResponseCode) {

	var errorfile string
	var stdoutfile string

	var iops_limit uint32 //最低基础限制，最低要大于等于这个
	var iops_max uint32
	var divide uint32
	var iops_goal uint32 //应该达到的性能

	var cmd string
	var cmderrorcheck string

	var errorcheckres string
	var iops_test_res uint32

	var all_test_disk_num uint32
	all_test_disk_num = 0

	switch disktype {
	case "RSSDDataDisk":

		iops_limit = 1800
		iops_max = FIO_RSSD_UDISK_IOPS_MAX
		divide = FIO_RSSD_DISKTYPE_IOPS_DIVIDE
	case "SSDDataDisk":

		iops_limit = 1200
		iops_max = FIO_SSD_UDISK_IOPS_MAX
		divide = FIO_SSD_DISKTYPE_IOPS_DIVIDE
	case "DataDisk":

		iops_limit = 1000 //暂定1000后面用到再减
		iops_max = FIO_DATA_UDISK_IOPS_MAX
		divide = 0

	default:
		ctx.ERRORF("No Such Udisk Type : %v | Please Check You Udisk Type", disktype)
		rc = NewResponseCode(FIO_TEST_DISKTYPE_ERROR)
		rc.SetAction(fmt.Sprintf("FioIopsTest/GetFioIopsTestResult"))
		return rc //处理错误
	}

	var hostTestUdiskAllNum map[string]uint32
	hostTestUdiskAllNum = make(map[string]uint32, 0)
	for _, v := range alluhostlist {
		tempInternalIp := v.Internal_Ip
		hostTestUdiskAllNum[tempInternalIp] = AllUhostTestDiskNum
	}

	for _, ipval := range iplist {
		hostTestUdiskAllNum[ipval] = (AllUhostTestDiskNum + iplistUhostAddTestDiskNum)
	}

	//DataDisk（普通数据盘），SSDDataDisk（SSD数据盘），RSSDDataDisk（RSSD数据盘）
	//GET_BW_WRITE_MODE_TEST_RESULT = " | grep WRITE | awk 'BEGIN{FS=\"=\"}{printf $4}' | tr -cd \"[0-9 .]\" |sed \"s/\\..*//g\" "
	if operation != "IopsStart" {   //以后添加直接获取测试结果的接口，直接在这里添加一个条件还有下面的if添加一个条件即可
		ctx.ERRORF("No Such operation In GetFioIopsTestResultByTestType func | You Operation Is  : %v | Please Check You Operation", operation)
		//处理错误
	}
	if operation == "IopsStart" {  //以后添加直接获取测试结果的接口，要这里添加一个条件
		for _, hostVal := range alluhostlist {
			ip_test_disknum := hostTestUdiskAllNum[hostVal.Internal_Ip]
			for i := uint32(0); i < ip_test_disknum; i++ {
				disk_nvme := hostVal.UDisk_App_Info[i].Disk_Lable
				disksize := fiocreatdisk.DiskSize //这是挂载盘大小，不是测试盘大小
				iops_test_res = 0
				all_test_disk_num += 1

				iops_goal = iops_limit + (disksize * divide)
				if divide == 0 && iops_limit == 1000 {
					iops_goal = iops_limit - 111
				}

				if iops_goal >= iops_max && divide != 0 {
					iops_goal = iops_max - 811
				}

				stdoutfile = "Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme + "_result.txt"
				errorfile = "Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme + "_error.txt"
				if fioParams.FioRWMode == "randread" {

					cmderrorcheck = "cat " + errorfile + " | wc -l"
					cmd = "cat " + stdoutfile + GET_IOPS_RANDREAD_TEST_CMD

				} else if fioParams.FioRWMode == "randwrite" {

					cmderrorcheck = "cat " + errorfile + " | wc -l"
					cmd = "cat " + stdoutfile + GET_IOPS_RANDWRITE_TEST_CMD

				} else if fioParams.FioRWMode == "randrw" {

					cmderrorcheck = "cat " + errorfile + " | wc -l"
					cmd = "cat " + stdoutfile + GET_IOPS_RANDRW_TEST_WRITE_CMD + "; cat " + stdoutfile + GET_IOPS_RANDRW_TEST_READ_CMD

				} else {
                          //处理error
					ctx.ERRORF("Fio Iops test No Such Mode, Please Check You Config Params [FioRWMode] | This Mode :[%v]  Is Error ", fioParams.FioRWMode)
					rc = NewResponseCode(FIO_TEST_RWMODE_ERROR)
				    rc.SetAction(fmt.Sprintf("FioIopsTest"))
					return rc 

				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "查看Fio IOPS测试错误日志是否有记录"
				OpTrack.CurrOpTk.RemoteOp.Ip = hostVal.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = cmderrorcheck
				OpTrack.CurrOpTk.RemoteOp.User = hostVal.User
				OpTrack.CurrOpTk.RemoteOp.Password = hostVal.Password

				errorcheckres, rc = RemoteCmdRequestWithResultWithRN(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Check Fio Iops Test Error Log Fail | Err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}

				OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestNoResult succeed")
				ctx.INFOF("Get Fio Iops Test Error File Data Number  Successful | The Error File Data Number Result Is : %v", errorcheckres)

				//去掉换行符，不然会报错，strconv.Atoi(diskNum) 为"0\n"
				errorcheckres = strings.Replace(errorcheckres, "\n", "", -1)

				if errorcheckres != "0" {
					ctx.ERRORF(" Fio Iops Test Error File Have Errors Or Get Error File Result Error  | Please Check errorTxt | You Check Error File Result Is : %v", errorcheckres)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "获取Fio IOPS测试结果"
				OpTrack.CurrOpTk.RemoteOp.Ip = hostVal.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = cmd
				OpTrack.CurrOpTk.RemoteOp.User = hostVal.User
				OpTrack.CurrOpTk.RemoteOp.Password = hostVal.Password

				var iops_res string
				iops_res, rc = RemoteCmdRequestWithResultWithRN(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Get Fio Iops Test Result  Fail | Err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}
				OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestNoResult succeed")
				ctx.INFOF("Get Fio Iops Test Result Successful | The Result Is : %v", iops_res)
				
				//去掉换行符，不然会报错，bw_res 为 ： "数字\n"

				iops_res = strings.Replace(iops_res, "\n", "", -1)
				iops_res = strings.Replace(iops_res, " ", "", -1) //去掉空格

				if fioParams.FioRWMode == "randrw" {
					temp_iops_res := strings.Split(iops_res, ",")
					if len(temp_iops_res) != FIO_IOPS_READ_WRITE_RESULT_LEN {
						ctx.ERRORF("Get Fio Iops Test Result Len Fail Maybe Fio Result Report Has Changeed| The Len Is %v : | Result Is : %v", len(temp_iops_res), temp_iops_res)
						rc = NewResponseCode(FIO_TEST_STRING_RES_CHANGE_TO_OTHER_TYPE_ERROR)
						rc.SetAction(fmt.Sprintf("FioIopsTest"))
						return rc 
						//处理错误
					}
					for _, iopsres := range temp_iops_res {
						temp_res, er := strconv.Atoi(iopsres)
						if er != nil {
						//处理错误
							ctx.ERRORF("Get Fio Iops Test Result Is String  | But String Type Change To Int Fail | Maybe String Result Error Or Fio Result Report Has Changeed | Err:%v", er)
							rc = NewResponseCode(FIO_TEST_STRING_RES_CHANGE_TO_OTHER_TYPE_ERROR)
							rc.SetAction(fmt.Sprintf("FioIopsTest"))
							return rc 
							
						}

						iops_test_res = iops_test_res + uint32(temp_res)
					}
					ctx.INFOF("输出fio的IOPS测试结果 ： Fio IOPS Test Result |All IOPS Is :  %v | Write Iops : %v  | Read Iops : %v  | 总共测试的盘数量: %v | 盘符:  %v   | 盘大小 : %v ", iops_test_res, temp_iops_res[0], temp_iops_res[1], all_test_disk_num, disk_nvme, disksize)

				} else {
					temp_res, er := strconv.Atoi(iops_res)
					if er != nil {
                      //处理错误
						ctx.ERRORF("Get Fio Iops Test Result Is String  | But String Type Change To Int Fail | Maybe String Result Error Or Fio Result Report Has Changeed | Err:%v", er)
						rc = NewResponseCode(FIO_TEST_STRING_RES_CHANGE_TO_OTHER_TYPE_ERROR)
				        rc.SetAction(fmt.Sprintf("FioIopsTest"))
						return rc 
						
					}
					iops_test_res = uint32(temp_res)
				}

				ctx.INFOF("输出fio的IOPS测试结果 : Fio IOPS Test Result | All IOPS Is :  %v    盘符:  %v   ,盘大小 : %v , 当前虚机 : %v 所有测试盘的数量 :%v", iops_test_res, disk_nvme, disksize, hostVal.Internal_Ip, all_test_disk_num)
				/*
					Header := GetEmailHeader(ctx.Email.UDiskSetId, HeaderInfo, CHAOS_ERROR_TITLE)
					err := SendEmail(Header, BodyInfo, ctx.Email.EmailReceiver)
				*/
			}
			ctx.INFOF("虚机 : %v  总共测试的盘数量: %v", hostVal.Internal_Ip, all_test_disk_num)
		}

	}

	ctx.INFOF("没有要进行fio测试的虚机，或者测试完成 you iperation is : %v , 所有的虚机总共测试盘数量 : %v", operation, all_test_disk_num)
	return
}




//备份    *****************************************************************


package methods

import (
	. "chaos/common/errs"
	. "chaos/common/exec"
	. "chaos/common/track"
	. "chaos/vm_controller/src/logic/context"
	. "chaos/vm_controller/src/logic/methods/methods_common"
	"fmt"
	"strconv"
	"strings"
)

type GETIPMODE = bool

var (
	fio_bs         string
	fio_iodepth    string
	fio_size       string
	fio_runtime    string
	fio_numjobs    string
	fio_filename   string
	fio_threadname string
	fio_rwmixread  string
)

const (
	//fio result const numbers
	FIO_ALL_DISKTYPE_BW_DISK_DIVIDE = uint32(2) //分配到磁盘的bw，就是1g分配1/2，这里写2，即可
	FIO_RSSD_UDISK_BW_MAX           = uint32(4800)
	FIO_SSD_UDISK_BW_MAX            = uint32(260)
	FIO_DATA_DISK_BW_MAX            = uint32(100)

	FIO_RSSD_UDISK_TIMEDELAY_MAX  = float64(0.2)
	FIO_RSSD_UDISK_TIMEDELAY_MINI = float64(0.1)
	FIO_SSD_UDISK_TIMEDELAY_MAX   = float64(3.0)
	FIO_SSD_UDISK_TIMEDELAY_MINI  = float64(0.5)
	FIO_DATA_UDISK_TIMEDELAY_MAX  = float64(10.0)
	FIO_DATA_UDISK_TIMEDELAY_MINI = float64(0.1)

	FIO_RSSD_DISKTYPE_IOPS_DIVIDE = uint32(50) //分配到磁盘的iops，rssd  1g分配 50
	FIO_SSD_DISKTYPE_IOPS_DIVIDE  = uint32(30)
	FIO_RSSD_UDISK_IOPS_MAX       = uint32(1200000)
	FIO_SSD_UDISK_IOPS_MAX        = uint32(24000)
	FIO_DATA_UDISK_IOPS_MAX       = uint32(1000)

	FIO_IOPS_READ_WRITE_RESULT_LEN = 2

	FIO_CMD_USERNAME = "root"
	NO_UHOST_INFO    = 0

	Fio_BW_Test_RWMode        = "read,write"
	Fio_TimeDelay_Test_RWMode = "randread,randwrite"

	GET_MOUNT_UDISK_NUMBER      = "lsblk | grep -v vda | grep -v NAME | wc -l"
	FIO_TEST_UDISK_NUMBER_ERROR = uint32(0)

	BW_READ_MODE_TEST_CMD  = "nohup fio -direct=1 -rw=read -ioengine=libaio -group_reporting -time_based=1 "
	BW_WRITE_MODE_TEST_CMD = "nohup fio -direct=1 -rw=write -ioengine=libaio -group_reporting -time_based=1 "

	GET_BW_WRITE_MODE_TEST_RESULT = " | grep WRITE | awk 'BEGIN{FS=\"=\"}{printf $4}' | tr -cd \"[0-9 .]\" |sed \"s/\\..*//g\" "
	GET_BW_READ_MODE_TEST_RESULT  = " | grep READ | awk 'BEGIN{FS=\"=\"}{printf $4}' | tr -cd \"[0-9 .]\" |sed \"s/\\..*//g\" "

	TIMEDELAY_RANDREAD_MODE_TEST_CMD  = "nohup fio -direct=1 -rw=randread -ioengine=libaio -group_reporting -time_based=1 "
	TIMEDELAY_RANDWRITE_MODE_TEST_CMD = "nohup fio -direct=1 -rw=randwrite -ioengine=libaio -group_reporting -time_based=1 "

	GET_TIMEDELAY_TEST_RESULT_UNIT_CMD = " | grep \"lat\" | grep \"avg\" | grep -v \"slat\" | grep -v \"clat\" | awk '{print $2}' | tr -d [:] | awk 'BEGIN{FS=\"(\"}{printf $2}' | awk 'BEGIN{FS=\")\"}{printf $1}';"
	GET_TIMEDELAY_TEST_RESULT_CMD      = " | grep \"lat\" | grep \"avg\" | grep -v \"slat\" | grep -v \"clat\" | awk '{print $5}' | awk 'BEGIN{FS=\"avg\"}{printf $2}' | awk 'BEGIN{FS=\",\"}{printf $1}'"

	IOPS_RANDREAD_MODE_TEST_CMD  = "nohup fio -direct=1 -rw=randread -ioengine=libaio -group_reporting -time_based=1 "
	IOPS_RANDWRITE_MODE_TEST_CMD = "nohup fio -direct=1 -rw=randwrite -ioengine=libaio -group_reporting -time_based=1 "
	IOPS_RWMIX_MODE_TEST_CMD     = "nohup fio -direct=1 -ioengine=libaio -group_reporting -time_based=1 -rw=randrw"

	GET_IOPS_RANDREAD_TEST_CMD     = " | grep iops | grep read | awk '{print $4}' | awk 'BEGIN{FS=\"=\"}{printf $2}' | awk 'BEGIN{FS=\",\"}{printf $1}'"
	GET_IOPS_RANDWRITE_TEST_CMD    = " | grep iops | grep write | awk '{print $4}' | awk 'BEGIN{FS=\"=\"}{printf $2}' | awk 'BEGIN{FS=\",\"}{printf $1}'"
	GET_IOPS_RANDRW_TEST_READ_CMD  = " | grep iops | grep read | awk 'BEGIN{FS=\"=\"}{printf $4}' | awk 'BEGIN{FS=\",\"}{printf $1}'"
	GET_IOPS_RANDRW_TEST_WRITE_CMD = " | grep iops | grep write | awk '{print $4}' | awk 'BEGIN{FS=\"=\"}{printf $2}'" //先write后read
	//GATE_IOPS_RANDRW_TEST_CMD = " | "
)

type DiskFioTestRes struct {
	Disk_Id        string
	Disk_Nvme      string
	Disk_Type      string
	Disk_Size      uint32
	Disk_RW_Mode   string
	Disk_BW_Res    uint32
	Disk_Iops_Res  uint32
	Disk_Delay_Res float32
}

type FioTestResult struct {
	All_Test_Disk_Num       uint32
	Achieve_Performance_Num uint32
	Test_Type               string
	Test_Host_ip            string    //虚机ip
	All_Disk_Res            []DiskFioTestRes
}

/*
var FioBWRWMode []string
FioBWRWMode = []string{"read","write"}
*/

//需要一个bool表示是否需要额外增加盘测试，不然不知道命令是否要额外增加盘测试，所以代码不好报错， 也可以不加，只要配置文件和使用方法详细介绍就可以了
func FioBWTest(ctx *InnerResourceInitContext, fioParams FioTestParams, alluhostlist []ApplyUHostInfo, iplist []string, AllUhostTestDiskNum uint32, iplistUhostAddTestDiskNum uint32, OpTrack *OperationTrack) (rc *ResponseCode) {

	ctx.ERRORF("进行Fio BW 压测获取的所有 uhost 详细信息是 : %v", alluhostlist)

	fio_bs = " -bs=" + fioParams.FioBs
	fio_iodepth = " -iodepth=" + fioParams.FioIodepth
	fio_size = " -size=" + fioParams.IoSize
	fio_runtime = " -runtime=" + fioParams.FioRuntime
	fio_numjobs = " -numjobs=" + fioParams.FioNumjobs
	fio_filename = " -filename=/dev/"

	var stdoutfile string
	var stderrorfile string
	var all_test_disk_num uint32
	
	if AllUhostTestDiskNum == FIO_TEST_UDISK_NUMBER_ERROR && iplistUhostAddTestDiskNum == FIO_TEST_UDISK_NUMBER_ERROR {

		ctx.ERRORF("You Fio Test Udisk Number Is Error, Please Check You Config | AllUhostTestDiskNum : %v  iplistUhostAddTestDiskNum  :  %v ", AllUhostTestDiskNum, iplistUhostAddTestDiskNum)
		rc = NewResponseCode(FIO_TEST_UDISK_NUM_IN_CONFIG_ERROR)
		rc.SetAction(fmt.Sprintf("FioBWTest"))
		return rc

	}
	all_test_disk_num = uint32(len(alluhostlist))*AllUhostTestDiskNum + uint32(len(iplist))*iplistUhostAddTestDiskNum

	var diskNum string
	var getUhostMountUdiskNum map[string]uint32
	getUhostMountUdiskNum = make(map[string]uint32, 0)

	OpTrack.CurrOpTk.IsRemoteOperation = true
	OpTrack.CurrOpTk.CurrentOperationId = 0

	for _, hostinfo := range alluhostlist {
		OpTrack.CurrOpTk.CurrentOperationDescrition = "获取虚机挂载udisk的个数"
		OpTrack.CurrOpTk.RemoteOp.Ip = hostinfo.Internal_Ip
		OpTrack.CurrOpTk.RemoteOp.Cmd = GET_MOUNT_UDISK_NUMBER
		OpTrack.CurrOpTk.RemoteOp.User = hostinfo.User
		OpTrack.CurrOpTk.RemoteOp.Password = hostinfo.Password

		diskNum, rc = RemoteCmdRequestWithResultWithRN(
			OpTrack.CurrOpTk.RemoteOp.User,
			OpTrack.CurrOpTk.RemoteOp.Password,
			OpTrack.CurrOpTk.RemoteOp.Ip,
			OpTrack.CurrOpTk.RemoteOp.Cmd)
		if rc != nil {
			ctx.ERRORF("Get uhost mount udisk number  fail | err:%v", rc)
			OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
			return rc
		}
		//去掉换行符，不然会报错，diskNum 为 ： "0\n"   ,strconv.Atoi(diskNum) 为0
		diskNum = strings.Replace(diskNum, "\n", "", -1)
		diskNumber, err := strconv.Atoi(diskNum)
		if err != nil {
			ctx.ERRORF("Get Fio BW Test Internal Ip Attach Udisk Number Is String  | But String Type Change To Int Fail | Maybe String Result Error Or Cmd Error | Err:%v", err)
			rc = NewResponseCode(FIO_TEST_STRING_RES_CHANGE_TO_OTHER_TYPE_ERROR)
			rc.SetAction(fmt.Sprintf("FioIopsTest"))
			return rc 
		
		}
		OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestWithResultWithRN succeed")
		getUhostMountUdiskNum[hostinfo.Internal_Ip] = uint32(diskNumber)
		ctx.INFOF("挂载的磁盘为 ：%v   uhost mount udisk number  | disk number %s", getUhostMountUdiskNum[hostinfo.Internal_Ip], diskNumber)
	}

	var allTestUdiskUhostList map[string]ApplyUHostInfo
	allTestUdiskUhostList = make(map[string]ApplyUHostInfo, 0)

	//注意 AllUhostTestDiskNum 可以进入这个函数一定不为空并且有虚机信息    iplistUhostAddTestDiskNumd对应的列表可能为nil

	for _, uhostinfo := range alluhostlist {
		allTestUdiskUhostList[uhostinfo.Internal_Ip] = uhostinfo
	}

	var addTestUdiskUhostList map[string]ApplyUHostInfo
	addTestUdiskUhostList = make(map[string]ApplyUHostInfo, 0)
	if len(iplist) > 0 && iplist[0] != "" {
		for _, ip := range iplist {
			value, ok := allTestUdiskUhostList[ip]
			if ok != true {
				ctx.ERRORF("addTestUdiskUhostList  |  ip : %v | not exit in allTestUdiskUhostList, please check you config or chaos mongo ", ip)
				rc = NewResponseCode(FIO_TEST_IP_NOT_EXIT)
				rc.SetAction(fmt.Sprintf("FioBWTest"))
				return rc
			}
			addTestUdiskUhostList[ip] = value

		}
	}

	var fio_cmd string
	//开始进行对所有的要测试fio的虚机的磁盘进行压测
	if AllUhostTestDiskNum != FIO_TEST_UDISK_NUMBER_ERROR {
		for _, uhostinfo := range alluhostlist {
			if getUhostMountUdiskNum[uhostinfo.Internal_Ip] < AllUhostTestDiskNum {
				//报错
				ctx.ERRORF("Fio test udisk number in uhost not enough maybe you config fio test udisk number too many | UhostMountUdiskNum : %v    AllUhostTestDiskNum  : %v", getUhostMountUdiskNum[uhostinfo.Internal_Ip], AllUhostTestDiskNum)
				rc = NewResponseCode(FIO_TEST_UDISK_NUM_NOT_ENOUGH)
				rc.SetAction(fmt.Sprintf("FioBWTest"))
				return rc
			}
			// 开始压测
			var testnum uint32
			for testnum = 0; testnum < AllUhostTestDiskNum; testnum++ {
				disk_nvme := uhostinfo.UDisk_App_Info[testnum].Disk_Lable

				stdoutfile = "Fio_BW_" + fioParams.FioRWMode + "_" + disk_nvme + "_result.txt"
				stderrorfile = "Fio_BW_" + fioParams.FioRWMode + "_" + disk_nvme + "_error.txt"
				fio_threadname = "  -name=Fio_BW_" + fioParams.FioRWMode + "_" + disk_nvme

				if fioParams.FioRWMode == "read" {

					fio_cmd = BW_READ_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else if fioParams.FioRWMode == "write" {

					fio_cmd = BW_WRITE_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else {

					ctx.ERRORF("Fio BW test No Such Mode, Please Check You Config Params [FioRWMode] | This Mode :[%v]  Is Error ", fioParams.FioRWMode)
					rc = NewResponseCode(FIO_TEST_RWMODE_ERROR)
				    rc.SetAction(fmt.Sprintf("FioBWTest"))
					return rc 
				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "开始对所有虚机进行fio压测"

				OpTrack.CurrOpTk.RemoteOp.User = uhostinfo.User
				OpTrack.CurrOpTk.RemoteOp.Password = uhostinfo.Password
				OpTrack.CurrOpTk.RemoteOp.Ip = uhostinfo.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = fio_cmd
				rc = RemoteCmdRequest(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Fio test start  fail, please check you commond or remot uhost  | err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestNoResult failed")
					return rc
				}

				OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestNoResult succeed")
				ctx.INFOF("Uhost : [%v]  Start Fio BWTest  Successful  This Uhost Test Udisk Nvme Is :%v",uhostinfo.Internal_Ip, disk_nvme)

				 // *****起动压测后看是否有错误*****
				 var errorcheckres string
				cmderrorcheck := "cat " + stderrorfile + " | wc -l"
				OpTrack.CurrOpTk.CurrentOperationDescrition = "查看Fio BW带宽测试错误日志是否有记录"
				OpTrack.CurrOpTk.RemoteOp.Ip = uhostinfo.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = cmderrorcheck
				OpTrack.CurrOpTk.RemoteOp.User = uhostinfo.User
				OpTrack.CurrOpTk.RemoteOp.Password = uhostinfo.Password

				errorcheckres, rc = RemoteCmdRequestWithResultWithRN(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Check Fio BW Test Error Log Failed  |  Error File Is : %v   Host Ip Is : %v  | Err:%v",stderrorfile ,uhostinfo.Internal_Ip, rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}
				OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestNoResult succeed")
				ctx.INFOF("Get Fio Iops Test Error File Data Number  Successful | The Error File Data Number Result Is : %v", errorcheckres)
				//OpTrack.DumpLog()

				//去掉换行符，不然会报错，strconv.Atoi(diskNum) 为"0\n"
				errorcheckres = strings.Replace(errorcheckres, "\n", "", -1)

				if errorcheckres != "0" {
					ctx.ERRORF(" Fio BW Test Error Txt Have Errors | Please Check errorTxt | Error File Is : %v   Host Ip Is : %v",stderrorfile ,uhostinfo.Internal_Ip)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}

			ctx.INFOF("Uhost : [%v]  Start Fio BWTest  Successful | Check The Error File Not Find ERROR")
			
			}

			ctx.INFOF("Uhost : [%v]  Start Fio AllHostList BWTest  Successful  This Uhost Test Udisk Number Is :%v",uhostinfo.Internal_Ip, testnum)

		}
		ctx.INFOF(" All Uhost Start Fio BWTest Successful | alluhostlist : [%v]   Per Uhost  Udisk Number : %v | Will Start iplistUhostAddTestDisk Fio Test",alluhostlist, uint32(len(alluhostlist))*AllUhostTestDiskNum)
	}
	//对需要额外加盘测试的虚机的磁盘进行fio压测    仔细考虑流程和各种情况
	if iplistUhostAddTestDiskNum != FIO_TEST_UDISK_NUMBER_ERROR && len(addTestUdiskUhostList) > 0 {
		for _, value := range addTestUdiskUhostList {

			if getUhostMountUdiskNum[value.Internal_Ip] < (iplistUhostAddTestDiskNum + AllUhostTestDiskNum) {
				//报错
				ctx.ERRORF("Fio test udisk number in uhost not enough , add not enough  maybe you config fio test udisk number too many | UhostMountUdiskNum : %v    AllUhostTestDiskNum  : %v     iplistUhostAddTestDiskNum  : %v", getUhostMountUdiskNum[value.Internal_Ip], AllUhostTestDiskNum, iplistUhostAddTestDiskNum)
				rc = NewResponseCode(FIO_TEST_UDISK_NUM_NOT_ENOUGH)
				rc.SetAction(fmt.Sprintf("FioBWTest"))
				return rc
			}
			var testnum uint32
			for testnum = AllUhostTestDiskNum; testnum < (iplistUhostAddTestDiskNum + AllUhostTestDiskNum); testnum++ {
				disk_nvme := value.UDisk_App_Info[testnum].Disk_Lable
				//2>>stderr.txt 1>stdout3.txt
				stdoutfile = "Fio_BW_" + fioParams.FioRWMode + "_" + disk_nvme + "_result.txt"
				stderrorfile = "Fio_BW_" + fioParams.FioRWMode + "_" + disk_nvme + "_error.txt"
				fio_threadname = "  -name=Fio_BW_" + fioParams.FioRWMode + "_" + disk_nvme

				if fioParams.FioRWMode == "read" {

					fio_cmd = BW_READ_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else if fioParams.FioRWMode == "write" {

					fio_cmd = BW_WRITE_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else {

					ctx.ERRORF("Fio BW test No Such Mode, Please Check You Config Params [FioRWMode] | This Mode :[%v]  Is Error ", fioParams.FioRWMode)
					rc = NewResponseCode(FIO_TEST_RWMODE_ERROR)
				    rc.SetAction(fmt.Sprintf("FioBWTest"))
					return rc 
				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "开始对虚机进行fio压测"

				OpTrack.CurrOpTk.RemoteOp.User = value.User
				OpTrack.CurrOpTk.RemoteOp.Password = value.Password
				OpTrack.CurrOpTk.RemoteOp.Ip = value.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = fio_cmd
				rc = RemoteCmdRequest(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Fio test start  fail, please check you commond or remot uhost  | err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestNoResult failed")
					return rc
				}
				OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestNoResult succeed")

				 // *****起动压测后看是否有错误*****
				var errorcheckres string
				cmderrorcheck := "cat " + stderrorfile + " | wc -l"
				OpTrack.CurrOpTk.CurrentOperationDescrition = "查看Fio BW带宽测试错误日志是否有记录"
				OpTrack.CurrOpTk.RemoteOp.Ip = value.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = cmderrorcheck
				OpTrack.CurrOpTk.RemoteOp.User = value.User
				OpTrack.CurrOpTk.RemoteOp.Password = value.Password

				errorcheckres, rc = RemoteCmdRequestWithResultWithRN(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Check Fio BW Test Error Log Fail |  Error File Is : %v   Host Ip Is : %v | Err:%v", rc ,stderrorfile ,value.Internal_Ip)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}
				OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestNoResult succeed")
				ctx.INFOF("Get Fio Iops Test Error File Data Number  Successful | The Error File Data Number Result Is : %v", errorcheckres)
	
				//去掉换行符，不然会报错，strconv.Atoi(diskNum) 为"0\n"
				errorcheckres = strings.Replace(errorcheckres, "\n", "", -1)

				if errorcheckres != "0" {
					ctx.ERRORF(" Fio BW Test Error Txt Have Errors | Please Check errorTxt | Error File Is : %v   Host Ip Is : %v",stderrorfile ,value.Internal_Ip)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}

			 ctx.INFOF("Uhost : [%v]  Start Fio BWTest  Successful | Check The Error File Not Find ERROR")

			}
			ctx.INFOF("Uhost : [%v]  Start Fio BWTest  Successful | This Uhost Test Udisk Number Is :%v",value.Internal_Ip, (testnum - AllUhostTestDiskNum))	

		}
		ctx.INFOF(" All iplistUhostAddTestDisk Start Fio BWTest Successful iplistUhostAddTestDisk : [%v]  |  Per Uhost  Udisk Number Is : %v",iplist, iplistUhostAddTestDiskNum)

	}

	ctx.INFOF("AllUhostTestDisk And iplistUhostAddTestDisk Start Fio BW Test Successful | All  Fio BWTest Udisk Number  is : %v ", all_test_disk_num)
	return nil

}





//通过配置文件的Iplist获取Iplist对应的Uhost信息 或者 直接查询数据库获取全部的Uhost的信息
func GetUhostByIpListOrDB(ctx *InnerResourceInitContext, action string, internalIpList []string, dbtag string, dbResourceType string, getFioTestIpMode GETIPMODE, OpTrack *OperationTrack) (uhostlist []ApplyUHostInfo, rc *ResponseCode) {

	var uhostlistinfo string
	if getFioTestIpMode == true { //通过配置文件的Iplist获取Iplist对应的Uhost信息

		OpTrack.CurrOpTk.IsRemoteOperation = false
		OpTrack.CurrOpTk.CurrentOperationId = 0
		OpTrack.CurrOpTk.CurrentOperationDescrition = "通过配置文件的 IPList 获取虚机"

		if len(internalIpList) < 0 || internalIpList[0] == "" {
			ctx.ERRORF("Get InternalIp By IpList Failed  |  Please Check You Config  |  You IpList Value Is : %v  |  IpList Len : %d", internalIpList, len(internalIpList))
			rc = NewResponseCode(GET_INTERNAL_IP_BY_CONFIG_ERROR)
			rc.SetAction(fmt.Sprintf(action))
			OpTrack.DumpLog(OPERATION_RESULT_FAIL, "Get InternalIp By IpList Failed Please Check You Config")
			return nil, rc

		}
		for _, oneip := range internalIpList {
			tempuhostinfo := make([]ApplyUHostInfo, 0)
			tempuhostinfo, rc = GetApplyUHostByInternalIp(oneip, dbResourceType)
			if rc != nil {
				ctx.ERRORF("Get InternalIp By IpList Failed Please Check You Config | You IpList : %v  dbResourceType : %s  | err:%v", internalIpList, dbResourceType, rc)
				rc = NewResponseCode(GET_UHOST_INFO_BY_MONGO_ERROR)
				rc.SetAction(fmt.Sprintf(action))
				return nil, rc
			}
			uhostlist = append(uhostlist, tempuhostinfo...)
		}
	} else { //直接查询数据库获取全部的Uhost的信息
		OpTrack.CurrOpTk.IsRemoteOperation = false
		OpTrack.CurrOpTk.CurrentOperationId = 0
		OpTrack.CurrOpTk.CurrentOperationDescrition = "通过 Chaos 的 DB 获取虚机信息"

		uhostlist, rc = GetApplyUHostByTag(dbtag, dbResourceType)
		if rc != nil {
			ctx.ERRORF("Get InternalIp By dbtag : %s And dbResourceType : %s  |   Failed Please Check You Config | err:%v", dbtag, dbResourceType, rc)
			rc = NewResponseCode(GET_UHOST_INFO_BY_MONGO_ERROR)
			rc.SetAction(fmt.Sprintf(action))
			return nil, rc
		}

	}

	for _, v := range uhostlist {
		tmpinfo := fmt.Sprintf("\nInternalIp : %s  ,  Uhost Id : %s\n", v.Internal_Ip, v.UHost_Id)
		uhostlistinfo = uhostlistinfo + tmpinfo

	}
	if len(uhostlist) <= NO_UHOST_INFO /*|| uhostlist[0] == "" */ {
		ctx.ERRORF("Get Uhost Info By[ dbtag : %s  or InternalIp : %v ] And dbResourceType : %s  |   Failed Please Check You Config or Chaos Mongo Params", dbtag, internalIpList, dbResourceType, rc)
		rc = NewResponseCode(GET_UHOST_INFO_BY_MONGO_ERROR)
		rc.SetAction(fmt.Sprintf(action))
		OpTrack.DumpLog(OPERATION_RESULT_FAIL, "Get Uhost Info Failed, All UHostInfo: "+uhostlistinfo+"UHostInfo Lenth : "+string(len(uhostlist)))
		return nil, rc
	}
	OpTrack.DumpLog(OPERATION_RESULT_SUCCEED, "Get Uhost Info Succeed, All UHostInfo: "+uhostlistinfo)
	rc = nil
	return uhostlist, rc

}

//BW  cmd :  cat
//                                                               集群的盘类型
func GetFioBWTestResultByTestType(ctx *InnerResourceInitContext, disktype string, fioParams FioTestParams, fiocreatdisk CreateUdiskParams, operation string, alluhostlist []ApplyUHostInfo, iplist []string, AllUhostTestDiskNum uint32, iplistUhostAddTestDiskNum uint32, OpTrack *OperationTrack) (rc *ResponseCode) {

	var errorfile string
	var stdoutfile string

	var bw_limit uint32 //最低基础限制
	var bw_max uint32
	var divide uint32
	var bw_goal uint32

	var cmd string
	var cmderrorcheck string

	var errorcheckres string
	var bw_test_res uint32

	var all_test_disk_num uint32
	all_test_disk_num = 0

	divide = FIO_ALL_DISKTYPE_BW_DISK_DIVIDE

	switch disktype {
	case "RSSDDataDisk":

		bw_limit = 120
		bw_max = FIO_RSSD_UDISK_BW_MAX
	case "SSDDataDisk":

		bw_limit = 80
		bw_max = FIO_SSD_UDISK_BW_MAX

	case "DataDisk":

		bw_limit = 100 //最低限制，暂时和追高一样，后面再减
		bw_max = FIO_DATA_DISK_BW_MAX
		divide = 0

	default:
		ctx.ERRORF("No Such Udisk Type : %v | Please Check You Udisk Type", disktype)
		return nil //处理错误
	}

	var hostTestUdiskAllNum map[string]uint32
	hostTestUdiskAllNum = make(map[string]uint32, 0)
	for _, v := range alluhostlist {
		tempInternalIp := v.Internal_Ip
		hostTestUdiskAllNum[tempInternalIp] = AllUhostTestDiskNum
	}

	for _, ipval := range iplist {
		hostTestUdiskAllNum[ipval] = (AllUhostTestDiskNum + iplistUhostAddTestDiskNum)
	}

	//DataDisk（普通数据盘），SSDDataDisk（SSD数据盘），RSSDDataDisk（RSSD数据盘）
	//GET_BW_WRITE_MODE_TEST_RESULT = " | grep WRITE | awk 'BEGIN{FS=\"=\"}{printf $4}' | tr -cd \"[0-9 .]\" |sed \"s/\\..*//g\" "
	if operation != "BWStart" {
		ctx.ERRORF("No Such operation In GetFioTestResultByTestType func | You Operation Is  : %v | Please Check You Operation", operation)
		//处理错误
	}
	if operation == "BWStart" {
		for _, hostVal := range alluhostlist {
			ip_test_disknum := hostTestUdiskAllNum[hostVal.Internal_Ip]
			for i := uint32(0); i < ip_test_disknum; i++ {
				disk_nvme := hostVal.UDisk_App_Info[i].Disk_Lable
				disksize := fiocreatdisk.DiskSize //这是挂载盘大小，不是测试盘大小
				all_test_disk_num += 1
				
				bw_goal = bw_limit + (disksize / divide)
				if divide == 0 && bw_limit == 100 {
					bw_goal = bw_limit - 11
				}
				if bw_goal >= bw_max && divide != 0 {
					bw_goal = bw_max - 8
				}

				stdoutfile = "Fio_BW_" + fioParams.FioRWMode + "_" + disk_nvme + "_result.txt"
				errorfile = "Fio_BW_" + fioParams.FioRWMode + "_" + disk_nvme + "_error.txt"
				if fioParams.FioRWMode == "read" {

					cmderrorcheck = "cat " + errorfile + " | wc -l"
					cmd = "cat " + stdoutfile + GET_BW_READ_MODE_TEST_RESULT

				} else if fioParams.FioRWMode == "write" {

					cmderrorcheck = "cat " + errorfile + " | wc -l"
					cmd = "cat " + stdoutfile + GET_BW_WRITE_MODE_TEST_RESULT

				} else {
					ctx.ERRORF("Fio BW Test No Such Test mode  : %v | Please Check You Test RW Type", fioParams.FioRWMode)
					return nil //处理error
				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "查看Fio BW带宽测试错误日志是否有记录"
				OpTrack.CurrOpTk.RemoteOp.Ip = hostVal.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = cmderrorcheck
				OpTrack.CurrOpTk.RemoteOp.User = hostVal.User
				OpTrack.CurrOpTk.RemoteOp.Password = hostVal.Password

				errorcheckres, rc = RemoteCmdRequestWithResultWithRN(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Check Fio BW Test Error Log Fail | Err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}
				//OpTrack.DumpLog()

				//去掉换行符，不然会报错，strconv.Atoi(diskNum) 为"0\n"
				errorcheckres = strings.Replace(errorcheckres, "\n", "", -1)

				if errorcheckres != "0" {
					ctx.ERRORF(" Fio BW Test Error Txt Have Errors | Please Check errorTxt")
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "获取Fio BW带宽测试结果"
				OpTrack.CurrOpTk.RemoteOp.Ip = hostVal.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = cmd
				OpTrack.CurrOpTk.RemoteOp.User = hostVal.User
				OpTrack.CurrOpTk.RemoteOp.Password = hostVal.Password

				var bw_res string
				bw_res, rc = RemoteCmdRequestWithResultWithRN(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Get Fio BW Test Result  Fail | Err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}
				//OpTrack.DumpLog()er
				//去掉换行符，不然会报错，bw_res 为 ： "数字\n"

				bw_res = strings.Replace(bw_res, "\n", "", -1)
				bw_res = strings.Replace(bw_res, " ", "", -1) //去掉空格
				temp_bw_res, er := strconv.Atoi(bw_res)
				if er != nil {
					ctx.ERRORF("Get Fio BW Test Result  Fail | Err:%v", er)

					//处理错误
				}
				bw_test_res = uint32(temp_bw_res) / 1024

				ctx.INFOF("输出fio的带宽测试结果 ： Fio BW Test Result | BW Is :  %vM/s    盘符:  %v   ,盘大小 : %v", bw_test_res, disk_nvme, disksize)
				/*
					Header := GetEmailHeader(ctx.Email.UDiskSetId, HeaderInfo, CHAOS_ERROR_TITLE)
					err := SendEmail(Header, BodyInfo, ctx.Email.EmailReceiver)
				*/
			}

		}

	}

	ctx.INFOF("没有要进行fio测试的虚机，或者测试完成 you iperation is : %v ", operation)
	return
}

func FioTimeDelayTest(ctx *InnerResourceInitContext, fioParams FioTestParams, alluhostlist []ApplyUHostInfo, iplist []string, AllUhostTestDiskNum uint32, iplistUhostAddTestDiskNum uint32, OpTrack *OperationTrack) (rc *ResponseCode) {

	ctx.ERRORF("获取的uhost详细信息是$$$$$$$$$$$$$$$$ : %v", alluhostlist)

	fio_bs = " -bs=" + fioParams.FioBs
	fio_iodepth = " -iodepth=" + fioParams.FioIodepth
	fio_size = " -size=" + fioParams.IoSize
	fio_runtime = " -runtime=" + fioParams.FioRuntime
	fio_numjobs = " -numjobs=" + fioParams.FioNumjobs
	fio_filename = " -filename=/dev/"

	var stdoutfile string
	var stderrorfile string
	var all_test_disk_num uint32

	if AllUhostTestDiskNum == FIO_TEST_UDISK_NUMBER_ERROR && iplistUhostAddTestDiskNum == FIO_TEST_UDISK_NUMBER_ERROR {

		ctx.ERRORF("You Fio Test Udisk Number Is Error, Please Check You Config |   AllUhostTestDiskNum : %v  iplistUhostAddTestDiskNum  :  %v ", AllUhostTestDiskNum, iplistUhostAddTestDiskNum)
		rc = NewResponseCode(FIO_TEST_UDISK_NUM_IN_CONFIG_ERROR)
		rc.SetAction(fmt.Sprintf("FioBWTest"))
		return rc

	}
	all_test_disk_num = uint32(len(alluhostlist))*AllUhostTestDiskNum + uint32(len(iplist))*iplistUhostAddTestDiskNum

	var diskNum string
	var getUhostMountUdiskNum map[string]uint32
	getUhostMountUdiskNum = make(map[string]uint32, 0)

	OpTrack.CurrOpTk.IsRemoteOperation = true
	OpTrack.CurrOpTk.CurrentOperationId = 0

	for _, hostinfo := range alluhostlist {
		OpTrack.CurrOpTk.CurrentOperationDescrition = "获取虚机挂载udisk的个数"
		OpTrack.CurrOpTk.RemoteOp.Ip = hostinfo.Internal_Ip
		OpTrack.CurrOpTk.RemoteOp.Cmd = GET_MOUNT_UDISK_NUMBER
		OpTrack.CurrOpTk.RemoteOp.User = hostinfo.User
		OpTrack.CurrOpTk.RemoteOp.Password = hostinfo.Password

		diskNum, rc = RemoteCmdRequestWithResultWithRN(
			OpTrack.CurrOpTk.RemoteOp.User,
			OpTrack.CurrOpTk.RemoteOp.Password,
			OpTrack.CurrOpTk.RemoteOp.Ip,
			OpTrack.CurrOpTk.RemoteOp.Cmd)
		if rc != nil {
			ctx.ERRORF("Get uhost mount udisk number  fail | err:%v", rc)
			OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
			return rc
		}
		//去掉换行符，不然会报错，diskNum 为 ： "0\n"   ,strconv.Atoi(diskNum) 为0
		diskNum = strings.Replace(diskNum, "\n", "", -1)
		diskNumber, err := strconv.Atoi(diskNum)
		if err != nil {
			fmt.Printf("666")
		}
		OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestWithResultWithRN succeed")
		getUhostMountUdiskNum[hostinfo.Internal_Ip] = uint32(diskNumber)
		ctx.ERRORF("挂载的磁盘为 ：%v   uhost mount udisk number  | disk number %s", getUhostMountUdiskNum[hostinfo.Internal_Ip], diskNumber)
	}

	var allTestUdiskUhostList map[string]ApplyUHostInfo
	allTestUdiskUhostList = make(map[string]ApplyUHostInfo, 0)

	//注意 AllUhostTestDiskNum 可以进入这个函数一定不为空并且有虚机信息    iplistUhostAddTestDiskNumd对应的列表可能为nil

	for _, uhostinfo := range alluhostlist {
		allTestUdiskUhostList[uhostinfo.Internal_Ip] = uhostinfo
	}

	var addTestUdiskUhostList map[string]ApplyUHostInfo
	addTestUdiskUhostList = make(map[string]ApplyUHostInfo, 0)
	if len(iplist) > 0 && iplist[0] != "" {
		for _, ip := range iplist {
			value, ok := allTestUdiskUhostList[ip]
			if ok != true {
				ctx.ERRORF("addTestUdiskUhostList  |  ip : %v | not exit in allTestUdiskUhostList, please check you config or chaos mongo ", ip)
				rc = NewResponseCode(FIO_TEST_IP_NOT_EXIT)
				rc.SetAction(fmt.Sprintf("FioBWTest"))
				return rc
			}
			addTestUdiskUhostList[ip] = value

		}
	}

	var fio_cmd string
	//开始进行对所有的要测试fio的虚机的磁盘进行压测
	if AllUhostTestDiskNum != FIO_TEST_UDISK_NUMBER_ERROR {
		for _, uhostinfo := range alluhostlist {
			if getUhostMountUdiskNum[uhostinfo.Internal_Ip] < AllUhostTestDiskNum {
				//报错
				ctx.ERRORF("Fio test udisk number in uhost not enough maybe you config fio test udisk number too many | UhostMountUdiskNum : %v    AllUhostTestDiskNum  : %v", getUhostMountUdiskNum[uhostinfo.Internal_Ip], AllUhostTestDiskNum)
				rc = NewResponseCode(FIO_TEST_UDISK_NUM_NOT_ENOUGH)
				rc.SetAction(fmt.Sprintf("FioBWTest"))
				return rc
			}
			// 开始压测
			var testnum uint32
			for testnum = 0; testnum < AllUhostTestDiskNum; testnum++ {
				disk_nvme := uhostinfo.UDisk_App_Info[testnum].Disk_Lable

				stdoutfile = "Fio_TimeDelay_" + fioParams.FioRWMode + "_" + disk_nvme + "_result.txt"
				stderrorfile = "Fio_TimeDelay_" + fioParams.FioRWMode + "_" + disk_nvme + "_error.txt"
				fio_threadname = "  -name=Fio_TimeDelay_" + fioParams.FioRWMode + "_" + disk_nvme

				if fioParams.FioRWMode == "randread" {
					fio_cmd = TIMEDELAY_RANDREAD_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"
				}
				if fioParams.FioRWMode == "randwrite" {
					fio_cmd = TIMEDELAY_RANDWRITE_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"
				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "开始对所有虚机进行fio压测"

				OpTrack.CurrOpTk.RemoteOp.User = uhostinfo.User
				OpTrack.CurrOpTk.RemoteOp.Password = uhostinfo.Password
				OpTrack.CurrOpTk.RemoteOp.Ip = uhostinfo.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = fio_cmd
				rc = RemoteCmdRequest(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Fio test start  fail, please check you commond or remot uhost  | err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestNoResult failed")
					return rc
				}
				OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestNoResult succeed")
			}

		}
	}
	//对需要额外加盘测试的虚机的磁盘进行fio压测    仔细考虑流程和各种情况
	if iplistUhostAddTestDiskNum != FIO_TEST_UDISK_NUMBER_ERROR && len(addTestUdiskUhostList) > 0 {
		for key, value := range addTestUdiskUhostList {

			if getUhostMountUdiskNum[value.Internal_Ip] < (iplistUhostAddTestDiskNum + AllUhostTestDiskNum) {
				//报错
				ctx.ERRORF("Fio test udisk number in uhost not enough , add not enough  maybe you config fio test udisk number too many | UhostMountUdiskNum : %v    AllUhostTestDiskNum  : %v     iplistUhostAddTestDiskNum  : %v", getUhostMountUdiskNum[value.Internal_Ip], AllUhostTestDiskNum, iplistUhostAddTestDiskNum)
				rc = NewResponseCode(FIO_TEST_UDISK_NUM_NOT_ENOUGH)
				rc.SetAction(fmt.Sprintf("FioBWTest"))
				return rc
				fmt.Printf(key)
			}
			var testnum uint32
			for testnum = AllUhostTestDiskNum; testnum < (iplistUhostAddTestDiskNum + AllUhostTestDiskNum); testnum++ {
				disk_nvme := value.UDisk_App_Info[testnum].Disk_Lable
				//2>>stderr.txt 1>stdout3.txt
				stdoutfile = "Fio_TimeDelay_" + fioParams.FioRWMode + "_" + disk_nvme + "_result.txt"
				stderrorfile = "Fio_TimeDelay_" + fioParams.FioRWMode + "_" + disk_nvme + "_error.txt"
				fio_threadname = "  -name=Fio_TimeDelay_" + fioParams.FioRWMode + "_" + disk_nvme

				if fioParams.FioRWMode == "randread" {

					fio_cmd = TIMEDELAY_RANDREAD_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else if fioParams.FioRWMode == "randwrite" {

					fio_cmd = TIMEDELAY_RANDWRITE_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else {
					ctx.ERRORF("Fio Time Delay test No Such Mode, Please Check You Config Params [FioRWMode] | This Mode :[%v]  Is Error ", fioParams.FioRWMode)
				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "开始对虚机进行fio压测"

				OpTrack.CurrOpTk.RemoteOp.User = value.User
				OpTrack.CurrOpTk.RemoteOp.Password = value.Password
				OpTrack.CurrOpTk.RemoteOp.Ip = value.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = fio_cmd
				rc = RemoteCmdRequest(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Fio test start  fail, please check you commond or remot uhost  | err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestNoResult failed")
					return rc
				}
				OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestNoResult succeed")
			}
		}

	}
	
	ctx.INFOF("Fio Test Udisk Number  is : %v ", all_test_disk_num)
	return nil

}

//                                                               集群的盘类型
func GetFioTimeDelayTestResultByTestType(ctx *InnerResourceInitContext, disktype string, fioParams FioTestParams, fiocreatdisk CreateUdiskParams, operation string, alluhostlist []ApplyUHostInfo, iplist []string, AllUhostTestDiskNum uint32, iplistUhostAddTestDiskNum uint32, OpTrack *OperationTrack) (rc *ResponseCode) {

	var errorfile string
	var stdoutfile string

	var timedelay_max float64
	var timedelay_min float64

	var cmd string
	var cmderrorcheck string

	var errorcheckres string
	var timedelay_test_res float64
	
	var all_test_disk_num uint32
	all_test_disk_num = 0

	switch disktype {
	case "RSSDDataDisk":

		timedelay_max = FIO_RSSD_UDISK_TIMEDELAY_MAX
		timedelay_min = FIO_SSD_UDISK_TIMEDELAY_MINI

	case "SSDDataDisk":

		timedelay_max = FIO_SSD_UDISK_TIMEDELAY_MAX
		timedelay_min = FIO_SSD_UDISK_TIMEDELAY_MINI

	case "DataDisk":

		timedelay_max = FIO_DATA_UDISK_TIMEDELAY_MAX
		timedelay_min = FIO_DATA_UDISK_TIMEDELAY_MINI

	default:
		ctx.ERRORF("No Such Udisk Type : %v | Please Check You Udisk Type", disktype)
		return nil //处理错误
	}

	var hostTestUdiskAllNum map[string]uint32
	hostTestUdiskAllNum = make(map[string]uint32, 0)
	for _, v := range alluhostlist {
		tempInternalIp := v.Internal_Ip
		hostTestUdiskAllNum[tempInternalIp] = AllUhostTestDiskNum
	}

	for _, ipval := range iplist {
		hostTestUdiskAllNum[ipval] = (AllUhostTestDiskNum + iplistUhostAddTestDiskNum)
	}

	//DataDisk（普通数据盘），SSDDataDisk（SSD数据盘），RSSDDataDisk（RSSD数据盘）
	//GET_BW_WRITE_MODE_TEST_RESULT = " | grep WRITE | awk 'BEGIN{FS=\"=\"}{printf $4}' | tr -cd \"[0-9 .]\" |sed \"s/\\..*//g\" "
	if operation != "TimeDelayStart" {
		ctx.ERRORF("No Such operation In GetFioTestResultByTestType func | You Operation Is  : %v | Please Check You Operation", operation)
		//处理错误
	}
	if operation == "TimeDelayStart" {
		for _, hostVal := range alluhostlist {
			ip_test_disknum := hostTestUdiskAllNum[hostVal.Internal_Ip]
			for i := uint32(0); i < ip_test_disknum; i++ {
				disk_nvme := hostVal.UDisk_App_Info[i].Disk_Lable
				disksize := fiocreatdisk.DiskSize //这是挂载盘大小，不是测试盘大小
				all_test_disk_num += 1

				stdoutfile = "Fio_TimeDelay_" + fioParams.FioRWMode + "_" + disk_nvme + "_result.txt"
				errorfile = "Fio_TimeDelay_" + fioParams.FioRWMode + "_" + disk_nvme + "_error.txt"
				if fioParams.FioRWMode == "randread" {

					cmderrorcheck = "cat " + errorfile + " | wc -l"
					cmd = "cat " + stdoutfile + GET_TIMEDELAY_TEST_RESULT_UNIT_CMD + "cat " + stdoutfile + GET_TIMEDELAY_TEST_RESULT_CMD

				} else if fioParams.FioRWMode == "randwrite" {

					cmderrorcheck = "cat " + errorfile + " | wc -l"
					cmd = "cat " + stdoutfile + GET_TIMEDELAY_TEST_RESULT_UNIT_CMD + "cat " + stdoutfile + GET_TIMEDELAY_TEST_RESULT_CMD

				} else {
					ctx.ERRORF("Fio BW Test No Such Test mode  : %v | Please Check You Test RW Type", fioParams.FioRWMode)
					return nil //处理error
				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "查看Fio IOPS测试错误日志是否有记录"
				OpTrack.CurrOpTk.RemoteOp.Ip = hostVal.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = cmderrorcheck
				OpTrack.CurrOpTk.RemoteOp.User = hostVal.User
				OpTrack.CurrOpTk.RemoteOp.Password = hostVal.Password

				errorcheckres, rc = RemoteCmdRequestWithResultWithRN(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Check Fio Iops Test Error Log Fail | Err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}
				//OpTrack.DumpLog()

				//去掉换行符，不然会报错，strconv.Atoi(diskNum) 为"0\n"
				errorcheckres = strings.Replace(errorcheckres, "\n", "", -1)

				if errorcheckres != "0" {
					ctx.ERRORF(" Fio BW Test Error Txt Have Errors | Please Check errorTxt")
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "获取Fio iops时延的测试结果"
				OpTrack.CurrOpTk.RemoteOp.Ip = hostVal.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = cmd
				OpTrack.CurrOpTk.RemoteOp.User = hostVal.User
				OpTrack.CurrOpTk.RemoteOp.Password = hostVal.Password

				var timedelay_res string
				timedelay_res, rc = RemoteCmdRequestWithResultWithRN(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Get Fio BW Test Result  Fail | Err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}
				//去掉换行符，不然会报错，bw_res 为 ： "数字\n"

				timedelay_res = strings.Replace(timedelay_res, "\n", "", -1)
				timedelay_res = strings.Replace(timedelay_res, " ", "", -1) //去掉空格
				sli := strings.Split(timedelay_res, "=")
				if len(sli) != 2 {
					ctx.ERRORF("Get TimeDelay Test Result  Fail | Maybe You Cmd Error or Fio Report Result Format Changed  | You Result Is :%v", sli)
					//获取时延失败
				}

				strtofloat, errres := strconv.ParseFloat(sli[1], 64)
				if errres != nil {

					ctx.ERRORF("TimeDelay Test Result String Change To Float64 Fail | Maybe You String : %s  Error or Other Errors  | Err:%v", sli[1], errres)

				}

				temp_timedelay_res, er := strconv.ParseFloat(fmt.Sprintf("%.2f", strtofloat), 64)
				if er != nil {

					ctx.ERRORF("TimeDelay Test Result Float64 Save Two Decimal Places Fail | Maybe You Parmas Error | You Params Should Is Float But The Parmas Is : [%T] : [%v]   Or Is Other Errors | Err:%v", strtofloat, strtofloat, er)
					//处理错误
				}

				if sli[0] == "usec" {

					timedelay_test_res = temp_timedelay_res / 1000

				} else if sli[0] == "msec" {

					timedelay_test_res = temp_timedelay_res

				} else {

					ctx.ERRORF("Get TimeDelay Test Result Time Unit Error | Maybe You Cmd Error or Fio Report Result Format Changed | You Can Look At The Fio Report Result Format | You Time Unit Is :%v", sli[0])

					//处理错误
				}

				ctx.INFOF("输出fio的时延测试结果 ： Fio TimeDelay Test Result | TimeDelay Is : %vmsec  盘符: %v | 盘大小 : %v | Max TimeDelay: %v | Min TimeDelay : %v  | 总共测试的盘数量 :%v", timedelay_test_res, disk_nvme, disksize, timedelay_max, timedelay_min, all_test_disk_num)
				/*
					Header := GetEmailHeader(ctx.Email.UDiskSetId, HeaderInfo, CHAOS_ERROR_TITLE)
					err := SendEmail(Header, BodyInfo, ctx.Email.EmailReceiver)
				*/
			}

		}

	}

	ctx.INFOF("没有要进行fio测试的虚机,或者测试完成,总共测试的盘数量 ：%v | you iperation is : %v ",all_test_disk_num, operation)
	return
}

func FioIopsTest(ctx *InnerResourceInitContext, fioParams FioTestParams, alluhostlist []ApplyUHostInfo, iplist []string, AllUhostTestDiskNum uint32, iplistUhostAddTestDiskNum uint32, OpTrack *OperationTrack) (rc *ResponseCode) {

	ctx.INFOF("获取的uhost详细信息是$$$$$$$$$$$$$$$$ : %v", alluhostlist)

	fio_bs = " -bs=" + fioParams.FioBs
	fio_iodepth = " -iodepth=" + fioParams.FioIodepth
	fio_size = " -size=" + fioParams.IoSize
	fio_runtime = " -runtime=" + fioParams.FioRuntime
	fio_numjobs = " -numjobs=" + fioParams.FioNumjobs
	fio_filename = " -filename=/dev/"
	fio_rwmixread = " -rwmixread=" + fioParams.FioIopsRWMixWrite

	var stdoutfile string
	var stderrorfile string

	var all_test_disk_num uint32

	if AllUhostTestDiskNum == FIO_TEST_UDISK_NUMBER_ERROR && iplistUhostAddTestDiskNum == FIO_TEST_UDISK_NUMBER_ERROR {

		ctx.ERRORF("You Fio Test Udisk Number Is Error, Please Check You Config |   AllUhostTestDiskNum : %v  iplistUhostAddTestDiskNum  :  %v ", AllUhostTestDiskNum, iplistUhostAddTestDiskNum)
		rc = NewResponseCode(FIO_TEST_UDISK_NUM_IN_CONFIG_ERROR)
		rc.SetAction(fmt.Sprintf("FioBWTest"))
		return rc

	}

	all_test_disk_num = uint32(len(alluhostlist))*AllUhostTestDiskNum + uint32(len(iplist))*iplistUhostAddTestDiskNum

	var diskNum string
	var getUhostMountUdiskNum map[string]uint32
	getUhostMountUdiskNum = make(map[string]uint32, 0)

	OpTrack.CurrOpTk.IsRemoteOperation = true
	OpTrack.CurrOpTk.CurrentOperationId = 0

	for _, hostinfo := range alluhostlist {
		OpTrack.CurrOpTk.CurrentOperationDescrition = "获取虚机挂载udisk的个数"
		OpTrack.CurrOpTk.RemoteOp.Ip = hostinfo.Internal_Ip
		OpTrack.CurrOpTk.RemoteOp.Cmd = GET_MOUNT_UDISK_NUMBER
		OpTrack.CurrOpTk.RemoteOp.User = hostinfo.User
		OpTrack.CurrOpTk.RemoteOp.Password = hostinfo.Password

		diskNum, rc = RemoteCmdRequestWithResultWithRN(
			OpTrack.CurrOpTk.RemoteOp.User,
			OpTrack.CurrOpTk.RemoteOp.Password,
			OpTrack.CurrOpTk.RemoteOp.Ip,
			OpTrack.CurrOpTk.RemoteOp.Cmd)
		if rc != nil {
			ctx.ERRORF("Get Uhost Mount Udisk Number  Fail In Ip : [%v] | err: %v", hostinfo.Internal_Ip, rc)
			OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
			return rc
		}
		//去掉换行符，不然会报错，diskNum 为 ： "0\n"   ,strconv.Atoi(diskNum) 为0
		diskNum = strings.Replace(diskNum, "\n", "", -1)
		diskNumber, err := strconv.Atoi(diskNum)
		if err != nil {
			fmt.Printf("666")
		}
		OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestWithResultWithRN succeed")
		getUhostMountUdiskNum[hostinfo.Internal_Ip] = uint32(diskNumber)
		ctx.INFOF("挂载磁盘的数量为 ：%v   uhost mount udisk number  | disk number %v", getUhostMountUdiskNum[hostinfo.Internal_Ip], diskNumber)
	}

	var allTestUdiskUhostList map[string]ApplyUHostInfo
	allTestUdiskUhostList = make(map[string]ApplyUHostInfo, 0)

	//注意 AllUhostTestDiskNum 可以进入这个函数一定不为空并且有虚机信息    iplistUhostAddTestDiskNumd对应的列表可能为nil

	for _, uhostinfo := range alluhostlist {
		allTestUdiskUhostList[uhostinfo.Internal_Ip] = uhostinfo
	}

	var addTestUdiskUhostList map[string]ApplyUHostInfo
	addTestUdiskUhostList = make(map[string]ApplyUHostInfo, 0)
	if len(iplist) > 0 && iplist[0] != "" {
		for _, ip := range iplist {
			value, ok := allTestUdiskUhostList[ip]
			if ok != true {
				ctx.ERRORF("addTestUdiskUhostList  |  ip : %v | not exit in allTestUdiskUhostList, please check you config or chaos mongo ", ip)
				rc = NewResponseCode(FIO_TEST_IP_NOT_EXIT)
				rc.SetAction(fmt.Sprintf("FioBWTest"))
				return rc
			}
			addTestUdiskUhostList[ip] = value
        ctx.INFOF("Fio Test Ip addTestUdiskUhostList Exit In allTestUdiskUhostList | You addTestUdiskUhostList One Ip Is : %v", ip)
		}
	}

	var fio_cmd string
	//开始进行对所有的要测试fio的虚机的磁盘进行压测
	if AllUhostTestDiskNum != FIO_TEST_UDISK_NUMBER_ERROR {
		for _, uhostinfo := range alluhostlist {
			if getUhostMountUdiskNum[uhostinfo.Internal_Ip] < AllUhostTestDiskNum {
				//报错
				ctx.ERRORF("Fio test udisk number in uhost not enough maybe you config fio test udisk number too many | UhostMountUdiskNum : %v    AllUhostTestDiskNum  : %v", getUhostMountUdiskNum[uhostinfo.Internal_Ip], AllUhostTestDiskNum)
				rc = NewResponseCode(FIO_TEST_UDISK_NUM_NOT_ENOUGH)
				rc.SetAction(fmt.Sprintf("FioIopsTest"))
				return rc
			}
			// 开始压测
			ctx.INFOF("Start Fio Test | Now  Test Ip Is : %v  Test Mode : %v ",uhostinfo.Internal_Ip, fioParams.FioRWMode)
			var testnum uint32
			for testnum = 0; testnum < AllUhostTestDiskNum; testnum++ {
				disk_nvme := uhostinfo.UDisk_App_Info[testnum].Disk_Lable

				stdoutfile = "Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme + "_result.txt"
				stderrorfile = "Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme + "_error.txt"
				fio_threadname = "  -name=Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme

				if fioParams.FioRWMode == "randread" {

					fio_cmd = IOPS_RANDREAD_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else if fioParams.FioRWMode == "randwrite" {

					fio_cmd = IOPS_RANDWRITE_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else if fioParams.FioRWMode == "randrw" {

					fio_cmd = IOPS_RWMIX_MODE_TEST_CMD + fio_rwmixread + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else {
					ctx.ERRORF("Fio Iops test No Such Mode, Please Check You Config Params [FioRWMode] | This Mode :[%v]  Is Error ", fioParams.FioRWMode)
					rc = NewResponseCode(FIO_TEST_RWMODE_ERROR)
				    rc.SetAction(fmt.Sprintf("FioIopsTest"))
					return rc 
				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "开始对所有虚机进行fio压测"

				OpTrack.CurrOpTk.RemoteOp.User = uhostinfo.User
				OpTrack.CurrOpTk.RemoteOp.Password = uhostinfo.Password
				OpTrack.CurrOpTk.RemoteOp.Ip = uhostinfo.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = fio_cmd
				rc = RemoteCmdRequest(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Fio test start fail, please check you commond or remot uhost  | err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestNoResult failed")
					return rc
				}
				OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestNoResult succeed")

				ctx.INFOF("Uhost : [%v]  Start Fio IopsTest  Successful  This Uhost Test Udisk Nvme Is :%v",uhostinfo.Internal_Ip, disk_nvme)
			}

			ctx.INFOF("Uhost : [%v]  Start Fio IopsTest  Successful  This Uhost Test Udisk Number Is :%v",uhostinfo.Internal_Ip, testnum)

		}
		ctx.INFOF(" All Uhost Start Fio IopsTest Successful | alluhostlist : [%v]   Per Uhost  Udisk Number : %v | Will Start iplistUhostAddTestDisk Fio Test",alluhostlist, uint32(len(alluhostlist))*AllUhostTestDiskNum)
	}
	//对需要额外加盘测试的虚机的磁盘进行fio压测    仔细考虑流程和各种情况
	if iplistUhostAddTestDiskNum != FIO_TEST_UDISK_NUMBER_ERROR && len(addTestUdiskUhostList) > 0 {
		for _, value := range addTestUdiskUhostList {

			if getUhostMountUdiskNum[value.Internal_Ip] < (iplistUhostAddTestDiskNum + AllUhostTestDiskNum) {
				//报错
				ctx.ERRORF("Fio test udisk number in uhost not enough , add not enough  maybe you config fio test udisk number too many | UhostMountUdiskNum : %v    AllUhostTestDiskNum  : %v     iplistUhostAddTestDiskNum  : %v", getUhostMountUdiskNum[value.Internal_Ip], AllUhostTestDiskNum, iplistUhostAddTestDiskNum)
				rc = NewResponseCode(FIO_TEST_UDISK_NUM_NOT_ENOUGH)
				rc.SetAction(fmt.Sprintf("FioBWTest"))
				return rc
			}
			var testnum uint32
			for testnum = AllUhostTestDiskNum; testnum < (iplistUhostAddTestDiskNum + AllUhostTestDiskNum); testnum++ {
				disk_nvme := value.UDisk_App_Info[testnum].Disk_Lable
				//2>>stderr.txt 1>stdout3.txt
				stdoutfile = "Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme + "_result.txt"
				stderrorfile = "Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme + "_error.txt"
				fio_threadname = "  -name=Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme

				if fioParams.FioRWMode == "randread" {

					fio_cmd = IOPS_RANDREAD_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else if fioParams.FioRWMode == "randwrite" {

					fio_cmd = IOPS_RANDWRITE_MODE_TEST_CMD + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else if fioParams.FioRWMode == "randrw" {

					fio_cmd = IOPS_RWMIX_MODE_TEST_CMD + fio_rwmixread + fio_bs + fio_iodepth + fio_size + fio_runtime + fio_numjobs + fio_filename + disk_nvme + fio_threadname + " 2>>" + stderrorfile + " 1>" + stdoutfile + " &"

				} else {
					ctx.ERRORF("Fio Iops test No Such Mode, Please Check You Config Params [FioRWMode] | This Mode :[%v]  Is Error ", fioParams.FioRWMode)
					rc = NewResponseCode(FIO_TEST_RWMODE_ERROR)
				    rc.SetAction(fmt.Sprintf("FioIopsTest"))
					return rc 
				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "开始对虚机进行fio压测"

				OpTrack.CurrOpTk.RemoteOp.User = value.User
				OpTrack.CurrOpTk.RemoteOp.Password = value.Password
				OpTrack.CurrOpTk.RemoteOp.Ip = value.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = fio_cmd
				rc = RemoteCmdRequest(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Fio test start  fail, please check you commond or remot uhost  | err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestNoResult failed")
					return rc
				}
				OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestNoResult succeed")
				ctx.INFOF("Uhost : [%v]  Start Fio IopsTest  Successful  This Uhost Test Udisk Nvme Is :%v",value.Internal_Ip, disk_nvme)
			}
			ctx.INFOF("Uhost : [%v]  Start Fio IopsTest  Successful | This Uhost Test Udisk Number Is :%v",value.Internal_Ip, testnum)
		}
		ctx.INFOF(" All iplistUhostAddTestDisk Start Fio IopsTest Successful iplistUhostAddTestDisk : [%v]  |  Per Uhost  Udisk Number : %v",iplist, uint32(len(iplist))*iplistUhostAddTestDiskNum)

	}

	ctx.INFOF("AllUhostTestDisk And iplistUhostAddTestDisk Start Fio Iops Test Successful | All  Fio Test Udisk Number  is : %v ", all_test_disk_num)
	return nil

}

func GetFioIopsTestResultByTestType(ctx *InnerResourceInitContext, disktype string, fioParams FioTestParams, fiocreatdisk CreateUdiskParams, operation string, alluhostlist []ApplyUHostInfo, iplist []string, AllUhostTestDiskNum uint32, iplistUhostAddTestDiskNum uint32, OpTrack *OperationTrack) (rc *ResponseCode) {

	var errorfile string
	var stdoutfile string

	var iops_limit uint32 //最低基础限制，最低要大于等于这个
	var iops_max uint32
	var divide uint32
	var iops_goal uint32 //应该达到的性能

	var cmd string
	var cmderrorcheck string

	var errorcheckres string
	var iops_test_res uint32

	var all_test_disk_num uint32
	all_test_disk_num = 0

	switch disktype {
	case "RSSDDataDisk":

		iops_limit = 1800
		iops_max = FIO_RSSD_UDISK_IOPS_MAX
		divide = FIO_RSSD_DISKTYPE_IOPS_DIVIDE
	case "SSDDataDisk":

		iops_limit = 1200
		iops_max = FIO_SSD_UDISK_IOPS_MAX
		divide = FIO_SSD_DISKTYPE_IOPS_DIVIDE
	case "DataDisk":

		iops_limit = 1000 //暂定1000后面用到再减
		iops_max = FIO_DATA_UDISK_IOPS_MAX
		divide = 0

	default:
		ctx.ERRORF("No Such Udisk Type : %v | Please Check You Udisk Type", disktype)
		rc = NewResponseCode(FIO_TEST_DISKTYPE_ERROR)
		rc.SetAction(fmt.Sprintf("FioIopsTest/GetFioIopsTestResult"))
		return rc //处理错误
	}

	var hostTestUdiskAllNum map[string]uint32
	hostTestUdiskAllNum = make(map[string]uint32, 0)
	for _, v := range alluhostlist {
		tempInternalIp := v.Internal_Ip
		hostTestUdiskAllNum[tempInternalIp] = AllUhostTestDiskNum
	}

	for _, ipval := range iplist {
		hostTestUdiskAllNum[ipval] = (AllUhostTestDiskNum + iplistUhostAddTestDiskNum)
	}

	//DataDisk（普通数据盘），SSDDataDisk（SSD数据盘），RSSDDataDisk（RSSD数据盘）
	//GET_BW_WRITE_MODE_TEST_RESULT = " | grep WRITE | awk 'BEGIN{FS=\"=\"}{printf $4}' | tr -cd \"[0-9 .]\" |sed \"s/\\..*//g\" "
	if operation != "IopsStart" {   //以后添加直接获取测试结果的接口，直接在这里添加一个条件还有下面的if添加一个条件即可
		ctx.ERRORF("No Such operation In GetFioIopsTestResultByTestType func | You Operation Is  : %v | Please Check You Operation", operation)
		//处理错误
	}
	if operation == "IopsStart" {  //以后添加直接获取测试结果的接口，要这里添加一个条件
		for _, hostVal := range alluhostlist {
			ip_test_disknum := hostTestUdiskAllNum[hostVal.Internal_Ip]
			for i := uint32(0); i < ip_test_disknum; i++ {
				disk_nvme := hostVal.UDisk_App_Info[i].Disk_Lable
				disksize := fiocreatdisk.DiskSize //这是挂载盘大小，不是测试盘大小
				iops_test_res = 0
				all_test_disk_num += 1

				iops_goal = iops_limit + (disksize * divide)
				if divide == 0 && iops_limit == 1000 {
					iops_goal = iops_limit - 111
				}

				if iops_goal >= iops_max && divide != 0 {
					iops_goal = iops_max - 811
				}

				stdoutfile = "Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme + "_result.txt"
				errorfile = "Fio_Iops_" + fioParams.FioRWMode + "_" + disk_nvme + "_error.txt"
				if fioParams.FioRWMode == "randread" {

					cmderrorcheck = "cat " + errorfile + " | wc -l"
					cmd = "cat " + stdoutfile + GET_IOPS_RANDREAD_TEST_CMD

				} else if fioParams.FioRWMode == "randwrite" {

					cmderrorcheck = "cat " + errorfile + " | wc -l"
					cmd = "cat " + stdoutfile + GET_IOPS_RANDWRITE_TEST_CMD

				} else if fioParams.FioRWMode == "randrw" {

					cmderrorcheck = "cat " + errorfile + " | wc -l"
					cmd = "cat " + stdoutfile + GET_IOPS_RANDRW_TEST_WRITE_CMD + "; cat " + stdoutfile + GET_IOPS_RANDRW_TEST_READ_CMD

				} else {
                          //处理error
					ctx.ERRORF("Fio Iops test No Such Mode, Please Check You Config Params [FioRWMode] | This Mode :[%v]  Is Error ", fioParams.FioRWMode)
					rc = NewResponseCode(FIO_TEST_RWMODE_ERROR)
				    rc.SetAction(fmt.Sprintf("FioIopsTest"))
					return rc 

				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "查看Fio IOPS测试错误日志是否有记录"
				OpTrack.CurrOpTk.RemoteOp.Ip = hostVal.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = cmderrorcheck
				OpTrack.CurrOpTk.RemoteOp.User = hostVal.User
				OpTrack.CurrOpTk.RemoteOp.Password = hostVal.Password

				errorcheckres, rc = RemoteCmdRequestWithResultWithRN(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Check Fio Iops Test Error Log Fail | Err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}

				OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestNoResult succeed")
				ctx.INFOF("Get Fio Iops Test Error File Data Number  Successful | The Error File Data Number Result Is : %v", errorcheckres)

				//去掉换行符，不然会报错，strconv.Atoi(diskNum) 为"0\n"
				errorcheckres = strings.Replace(errorcheckres, "\n", "", -1)

				if errorcheckres != "0" {
					ctx.ERRORF(" Fio Iops Test Error File Have Errors Or Get Error File Result Error  | Please Check errorTxt | You Check Error File Result Is : %v", errorcheckres)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}

				OpTrack.CurrOpTk.CurrentOperationDescrition = "获取Fio IOPS测试结果"
				OpTrack.CurrOpTk.RemoteOp.Ip = hostVal.Internal_Ip
				OpTrack.CurrOpTk.RemoteOp.Cmd = cmd
				OpTrack.CurrOpTk.RemoteOp.User = hostVal.User
				OpTrack.CurrOpTk.RemoteOp.Password = hostVal.Password

				var iops_res string
				iops_res, rc = RemoteCmdRequestWithResultWithRN(
					OpTrack.CurrOpTk.RemoteOp.User,
					OpTrack.CurrOpTk.RemoteOp.Password,
					OpTrack.CurrOpTk.RemoteOp.Ip,
					OpTrack.CurrOpTk.RemoteOp.Cmd)
				if rc != nil {
					ctx.ERRORF("Get Fio Iops Test Result  Fail | Err:%v", rc)
					OpTrack.DumpLog(OPERATION_RESULT_FAIL, "RemoteCmdRequestWithResultWithRN failed")
					return rc
				}
				OpTrack.DumpMongoDB(OPERATION_RESULT_SUCCEED, "RemoteCmdRequestNoResult succeed")
				ctx.INFOF("Get Fio Iops Test Result Successful | The Result Is : %v", iops_res)
				
				//去掉换行符，不然会报错，bw_res 为 ： "数字\n"

				iops_res = strings.Replace(iops_res, "\n", "", -1)
				iops_res = strings.Replace(iops_res, " ", "", -1) //去掉空格

				if fioParams.FioRWMode == "randrw" {
					temp_iops_res := strings.Split(iops_res, ",")
					if len(temp_iops_res) != FIO_IOPS_READ_WRITE_RESULT_LEN {
						ctx.ERRORF("Get Fio Iops Test Result Len Fail Maybe Fio Result Report Has Changeed| The Len Is %v : | Result Is : %v", len(temp_iops_res), temp_iops_res)
						rc = NewResponseCode(FIO_TEST_STRING_RES_CHANGE_TO_OTHER_TYPE_ERROR)
						rc.SetAction(fmt.Sprintf("FioIopsTest"))
						return rc 
						//处理错误
					}
					for _, iopsres := range temp_iops_res {
						temp_res, er := strconv.Atoi(iopsres)
						if er != nil {
						//处理错误
							ctx.ERRORF("Get Fio Iops Test Result Is String  | But String Type Change To Int Fail | Maybe String Result Error Or Fio Result Report Has Changeed | Err:%v", er)
							rc = NewResponseCode(FIO_TEST_STRING_RES_CHANGE_TO_OTHER_TYPE_ERROR)
							rc.SetAction(fmt.Sprintf("FioIopsTest"))
							return rc 
							
						}

						iops_test_res = iops_test_res + uint32(temp_res)
					}
					ctx.INFOF("输出fio的IOPS测试结果 ： Fio IOPS Test Result |All IOPS Is :  %v | Write Iops : %v  | Read Iops : %v  | 总共测试的盘数量: %v | 盘符:  %v   | 盘大小 : %v ", iops_test_res, temp_iops_res[0], temp_iops_res[1], all_test_disk_num, disk_nvme, disksize)

				} else {
					temp_res, er := strconv.Atoi(iops_res)
					if er != nil {
                      //处理错误
						ctx.ERRORF("Get Fio Iops Test Result Is String  | But String Type Change To Int Fail | Maybe String Result Error Or Fio Result Report Has Changeed | Err:%v", er)
						rc = NewResponseCode(FIO_TEST_STRING_RES_CHANGE_TO_OTHER_TYPE_ERROR)
				        rc.SetAction(fmt.Sprintf("FioIopsTest"))
						return rc 
						
					}
					iops_test_res = uint32(temp_res)
				}

				ctx.INFOF("输出fio的IOPS测试结果 : Fio IOPS Test Result | All IOPS Is :  %v    盘符:  %v   ,盘大小 : %v , 当前虚机 : %v 所有测试盘的数量 :%v", iops_test_res, disk_nvme, disksize, hostVal.Internal_Ip, all_test_disk_num)
				/*
					Header := GetEmailHeader(ctx.Email.UDiskSetId, HeaderInfo, CHAOS_ERROR_TITLE)
					err := SendEmail(Header, BodyInfo, ctx.Email.EmailReceiver)
				*/
			}
			ctx.INFOF("虚机 : %v  总共测试的盘数量: %v", hostVal.Internal_Ip, all_test_disk_num)
		}

	}

	ctx.INFOF("没有要进行fio测试的虚机，或者测试完成 you iperation is : %v , 所有的虚机总共测试盘数量 : %v", operation, all_test_disk_num)
	return
}

//2.每个虚机测试几块盘，指定具体的虚机多少盘

//相当于进制的十位(几个ip就是几进制)              个位
// 全部ip的都要挂一块盘                  1 2 3 4挂第二块盘(要在全部的范围内)   其余的不挂第二块盘
//结果存放文件名字 : BW_vdb.txt         错误输出的文件 ： BW_vdb_error.txt

//开始测试   for  range  ip

//  lsblk	获取盘符

//  开始测试是取ip里的盘符 ，但是要和lsblk里对比

/*



// 第几次压测         虚机个数         每个虚机的盘个数     总共盘个数             是否达到性能                            参数
//    1                 12                    1              12                     是                  虚机信息 每个虚机压测多少盘 [1]（针对虚机信息所有虚机）         [无]  每个虚机压测多少盘（针对虚机信息部分虚机，这里通过   传参设置虚机信息或者配置信息里配置虚机   那几台虚机压测几块盘）
//    2                 12                    2              24                     否                              [2]                                                    [无]
//    3                 12                    1              12 + n                 是                               1                                                       n
//    4                 12                    1              12 + n + m             否                               1                                                      n + m
//    5                 12                    1              12 + n + m - 1         是                               1                                                     n + m -1

*/







