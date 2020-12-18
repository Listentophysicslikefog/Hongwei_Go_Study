package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
	"sort"
	"strconv"
	"crypto/sha1"

)

type createUDiskResponse struct {
	ResponseCode
	UDiskId []string
}
	func PrettyPrintu(i interface{}) string {
		s, _ := json.MarshalIndent(i, "", "\t")
		return string(s)
	
	}

	type ResponseCode struct {
		RetCode int    `json:"RetCode"`
		Message string `json:"Message"`
		Action  string `json:"Action"` // 只有返回给api调用者时才会填充
	}

	func CreateUDisk() (udisk_ids []string, rc *ResponseCode) {
		params := make(map[string]interface{}, 0)
		params["Action"] = "CreateUDisk"
		params["request_uuid"] = "hongwei_request_test"
		params["Region"] = "cn-gd2"
		params["Zone"] = "cn-gd2-01"
		params["ProjectId"] = "org-nf3wlo"
		params["PublicKey"] = "O9aQkG61gppmkiPhiZyuljs+2yPwcmSXRMX9wI7yZujZ/qlOPKpPew=="
		params["Size"] = uint32(20)
		params["Name"] = "qika20201210"
		params["ChargeType"] = "Month"
		params["Quantity"] = int(1)
		params["UDataArkMode"] = "No"
		params["Tag"] = "qika_20201210_hn08_5009"   //"ChaosEris_hongwei.liu_20200727_test"
		params["DiskType"] = "RSSDDataDisk"
		params["UKmsMode"] = "No"
		params["SetId"] = uint32(5009)
	
		var rsp createUDiskResponse
		 HTTPRequdisk(params, &rsp)
	   time.Sleep(time.Second*20)
		fmt.Printf("create udisk result, %+v", PrettyPrintu(rsp))
		udisk_ids = rsp.UDiskId
		fmt.Println(udisk_ids)
		return
	}
	
	func CreateUDisks() (udisk_ids []string, rc *ResponseCode) {
		params := make(map[string]interface{}, 0)
		params["Action"] = "CreateUDisk"
		params["request_uuid"] = "hongwei_request_test"
		params["Region"] = "cn-gd2"
		params["Zone"] = "cn-gd2-01"
		params["ProjectId"] = "org-nf3wlo"
		params["PublicKey"] = "O9aQkG61gppmkiPhiZyuljs+2yPwcmSXRMX9wI7yZujZ/qlOPKpPew=="
		params["Size"] = uint32(20)
		params["Name"] = "qika20201210"
		params["ChargeType"] = "Month"
		params["Quantity"] = int(1)
		params["UDataArkMode"] = "No"
		params["Tag"] = "qika_20201210_hn08_5009"   //"ChaosEris_hongwei.liu_20200727_test"
		params["DiskType"] = "RSSDDataDisk"
		params["UKmsMode"] = "No"
		params["SetId"] = uint32(5009)
	
		var rsp createUDiskResponse
		 HTTPRequdisk(params, &rsp)
	   time.Sleep(time.Second*20)
		fmt.Printf("create udisk result, %+v", PrettyPrintu(rsp))
		udisk_ids = rsp.UDiskId
		fmt.Println(udisk_ids)
		return
	}

	func HTTPRequdisk(params map[string]interface{}, rsp interface{}) {
		Signature := Verfy_AC(params)
		params["Signature"] = Signature
		externalUrl := "https://api.upc-poc.cn"
		fmt.Printf("Post url: %s, params:%v", externalUrl, PrettyPrintu(params))
		b, err := json.Marshal(params)
		if err != nil {
			fmt.Printf("json err: %v", err)
			return
		}
		body := []byte(b)
		rspData, rc := SendHttpPostJsonRequestudisk(externalUrl, body)
		if rc != nil {
			fmt.Printf("http post json request fail; params:%v; rc:%v", params, rc)
			return
		}
		if err := json.Unmarshal(rspData, rsp); err != nil {
			fmt.Printf("unmarshal http rsp fail; err:%v", err)
			return
		}
	
		vRsp := reflect.ValueOf(rsp).Elem()
		retCode := vRsp.FieldByName("RetCode").Int()
		if retCode != 0 {
			retMessage := vRsp.FieldByName("Message").String()
			fmt.Printf("http  error; retcode:%d; message:%v;", retCode, retMessage)
			return
		}
	
		return
	}

	func Verfy_AC(params map[string]interface{}) (Signature string) {
		var SortString []string
		var ParamsData string = ""
		for k, _ := range params {
			SortString = append(SortString, k)
		}
		sort.Strings(SortString)
		for _, k := range SortString {
			fmt.Printf("Key = %s, Value = %v \n", k, params[k])
			ParamsData = ParamsData + k
			if str, ok := params[k].(string); ok {
				ParamsData += str
			} else if Param_i, ok := params[k].(int); ok {
				strParam := strconv.Itoa(Param_i)
				ParamsData += strParam
			} else if Param_ui32, ok := params[k].(uint32); ok {
				Param_ui64 := uint64(Param_ui32)
				strParam := strconv.FormatUint(Param_ui64, 10)
				ParamsData += strParam
			} else if Param_ui64, ok := params[k].(uint64); ok {
				strParam := strconv.FormatUint(Param_ui64, 10)
				ParamsData += strParam
			} else {
				
				panic("params type is invalid, current just accept string and int")
			}
		}
		ParamsData += "rLFMmP03h8xQCDbkq0RTphuiOfsgUabDVWcqm6o35JvxWC6mZZvsQL6244+VT51r"
		h := sha1.New()
		h.Write([]byte(ParamsData))
		bs := h.Sum(nil) // bs 是 16 进制的hash值
		Signature = fmt.Sprintf("%x", bs)
		return Signature
	}

	func SendHttpPostJsonRequestudisk(url string, req_body []byte) (rsp_data []byte, rc *ResponseCode) {
		rsp, err := http.Post(url, "application/json", bytes.NewBuffer(req_body))
		if err != nil {
			fmt.Printf("http json PostForm fail;err:%s", err)
			return
		}
	
		defer rsp.Body.Close()
		if rsp.StatusCode != 200 {
			fmt.Printf("http json PostForm rsp not OK;status_code:%d", rsp.StatusCode)
			fmt.Printf("rsp Header :%v", rsp.Header)
			body, _ := ioutil.ReadAll(rsp.Body)
			fmt.Printf("rsp body :%s", string(body))
			return
		}
		rsp_data, err = ioutil.ReadAll(rsp.Body)
	
		if err != nil {
			fmt.Printf("read json PostForm rsp fail;err:%s", err)
			return
		}
		return
	}

func main(){
		var rc *ResponseCode

		var udisk_ids []string
		var udisk_id []string
		//udisk_ids,rc=CreateUDisk()
for i:= 0 ; i< 1 ;i++ {
	udisk_ids,rc=CreateUDisks()
	if rc!= nil{
		fmt.Println("error")
		return
	}
	udisk_id = append(udisk_id,udisk_ids[0])
    }
		
		fmt.Println(udisk_id)
}




/*
Key = Action, Value = CreateUDisk
Key = ChargeType, Value = Month 
Key = DiskType, Value = RSSDDataDisk 
Key = Name, Value = hongwei0728 
Key = ProjectId, Value = org-nf3wlo 
Key = PublicKey, Value = O9aQkG61gppmkiPhiZyuljs+2yPwcmSXRMX9wI7yZujZ/qlOPKpPew== 
Key = Quantity, Value = 1 
Key = Region, Value = cn-gd2 
Key = SetId, Value = 5012 
Key = Size, Value = 300 
Key = Tag, Value = ChaosEris_hongwei.liu_20200727_test 
Key = UDataArkMode, Value = No 
Key = UKmsMode, Value = No 
Key = Zone, Value = cn-gd2-01 
Key = request_uuid, Value = hongwei_request
Post url: https://api.upc-poc.cn, params:{
	"Action": "CreateUDisk",
	"ChargeType": "Month",
	"DiskType": "RSSDDataDisk",
	"Name": "hongwei0728",
	"ProjectId": "org-nf3wlo",
	"PublicKey": "O9aQkG61gppmkiPhiZyuljs+2yPwcmSXRMX9wI7yZujZ/qlOPKpPew==",
	"Quantity": 1,
	"Region": "cn-gd2",
	"SetId": 5012,
	"Signature": "a9358098407a1eb736a56b02b2bbbde6927e82e7",
	"Size": 300,
	"Tag": "ChaosEris_hongwei.liu_20200727_test",
	"UDataArkMode": "No",
	"UKmsMode": "No",
	"Zone": "cn-gd2-01",
	"request_uuid": "hongwei_request"
}create udisk result, {
	"RetCode": 0,
	"Message": "",
	"Action": "CreateUDiskResponse",
	"UDiskId": [
		"bsr-f5xubrp3"
	]
}[bsr-f5xubrp3] <nil>

*/

/*
create udisk result, {
	"RetCode": 0,
	"Message": "",
	"Action": "CreateUDiskResponse",
	"UDiskId": [
		"bsr-xyndy4wo"
	]
}[bsr-xyndy4wo] <nil>

*/





/*0806
Key = Name, Value = hongwei0806 
Key = ProjectId, Value = org-nf3wlo 
Key = PublicKey, Value = O9aQkG61gppmkiPhiZyuljs+2yPwcmSXRMX9wI7yZujZ/qlOPKpPew== 
Key = Quantity, Value = 1 
Key = Region, Value = cn-gd2 
Key = SetId, Value = 5012 
Key = Size, Value = 300 
Key = Tag, Value = ChaosEris_hongwei.liu_20200727_test 
Key = UDataArkMode, Value = No 
Key = UKmsMode, Value = No 
Key = Zone, Value = cn-gd2-01 
Key = request_uuid, Value = hongwei_request 
Post url: https://api.upc-poc.cn, params:{
	"Action": "CreateUDisk",
	"ChargeType": "Month",
	"DiskType": "RSSDDataDisk",
	"Name": "hongwei0806",
	"ProjectId": "org-nf3wlo",
	"PublicKey": "O9aQkG61gppmkiPhiZyuljs+2yPwcmSXRMX9wI7yZujZ/qlOPKpPew==",
	"Quantity": 1,
	"Region": "cn-gd2",
	"SetId": 5012,
	"Signature": "a33e66e12c516d76707a7341a4e87eca0e30938e",
	"Size": 300,
	"Tag": "ChaosEris_hongwei.liu_20200727_test",
	"UDataArkMode": "No",
	"UKmsMode": "No",
	"Zone": "cn-gd2-01",
	"request_uuid": "hongwei_request"
}create udisk result, {
	"RetCode": 0,
	"Message": "",
	"Action": "CreateUDiskResponse",
	"UDiskId": [
		"bsr-yp0brqvn"
	]
}[bsr-yp0brqvn] <nil>


*/
