package main
import (
	"fmt"
)


type uhostfioinfo  struct {
     nowrunfionvme         map[string]bool 
	//nowrunfionvme         []map[string]bool                 //当前跑的盘符以及是否达到性能
    uhostinfo             ApplyUHostInfo                 
	//nowrunfionvme         []string                         //用来跑fio的盘符，uhost里的是所有的盘符 ,第一次全部的虚机的uhostfioinfo结构体里用来跑fio的盘符都设置为nvb
	currentlastnvme         string       
	achievePerformance      bool                               //分析结果后更改状态，初始为ture，如果结果达到性能，用来跑fio的盘符增加一个，如果不满足，盘符不变

}

var uhostinfo map[string]uhostfioinfo     //key  ip    value  hhostinfo

type UDisk_App_Pair struct {
	Disk_Lable               string           //盘符
	Attached_DataUDisk_Id    string           //挂载的数据盘extern_Id
	Run_Service              string           //udisk用途
	Run_Fio                  bool             //是否运行fio
}

type ApplyUHostInfo struct {
	UHost_Id              string
	UHost_Name            string
	Pysical_Mathine_Ip    string
	Internal_Ip           string
	User                  string
	Password              string
	UDisk_App_Info        []UDisk_App_Pair
	Have_Deploied_Fio     bool
}

type FioUHostInfo struct {
	UHostInfo               ApplyUHostInfo
	nowrunfionvme           map[string]bool 
	currentlastnvme         string       
	achievePerformance      bool     
}

var uhost_list = make([]FioUHostInfo, 0)
uhost_list, rc = GetApplyUHostByRunFio(ALL_INTERNALIP_WILL_RUN_FIO)

for k, v := range uhost_list{
var res  FioUHostInfo


	
}

func main(){




}

udisk_ids, rc = CreateUDisk(&ctx.InnerCtx, ctx.UdiskSetId, ctx.UdiskSize, udisk_name,
	UDiskChargeType, UDiskQuantity, UDiskUDataArkMode, ctx.Tag, ctx.DiskType,
	UDiskUKmsMode, UDiskCmkId, UDiskCouponId)

	func CreateUDisk(ctx *InnerResourceInitContext, udisk_setid uint32, Size uint32, Name string,
		ChargeType string, Quantity int, UDataArkMode string, Tag string, DiskType string,
		UKmsMode string, CmkId string, CouponId string) (udisk_ids []string, rc *ResponseCode) {
		params := make(map[string]interface{}, 0)
		params["Action"] = "CreateUDisk"
		params["request_uuid"] = ctx.Session
		params["ChildSession"] = ctx.ChildSession
		params["Region"] = ctx.Region
		params["Zone"] = ctx.Zone
		params["ProjectId"] = ctx.ProjectId
		params["PublicKey"] = ctx.PublicKey
		params["Size"] = Size
		params["Name"] = Name
		params["ChargeType"] = ChargeType
		params["Quantity"] = Quantity
		params["UDataArkMode"] = UDataArkMode
		params["Tag"] = Tag
		params["DiskType"] = DiskType
		params["UKmsMode"] = UKmsMode
		params["SetId"] = udisk_setid
	
		var rsp createUDiskResponse
		rc = HTTPReq(ctx, params, &rsp)
		if rc != nil {
			return
		}
	
		ctx.DEBUGF("create udisk result, %+v", PrettyPrint(rsp))
		udisk_ids = rsp.UDiskId
		return
	}
	








































