package main
import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"time"
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)

}


type createUHostInstanceResponse struct {
	ResponseCode
	UHostIds []string
	IPs      []string
}
type ResponseCode struct {
	RetCode int    `json:"RetCode"`
	Message string `json:"Message"`
	Action  string `json:"Action"` // 只有返回给api调用者时才会填充
}
//创建UHost实例
func CreateUHostInstance() (uhost_ids []string, ips []string) {
	params := make(map[string]interface{}, 0)
	params["Action"] = "CreateUHostInstance"
	params["request_uuid"] = "hongwei_request"
	params["Region"] = "cn-gd2"
	params["Zone"] = "cn-gd2-01"
	params["ProjectId"] = "org-nf3wlo"
	params["PublicKey"] = "O9aQkG61gppmkiPhiZyuljs+2yPwcmSXRMX9wI7yZujZ/qlOPKpPew=="
	params["ImageId"] = "uimage-rtwygh"
	params["Password"] = "UGFzc3dvcmQx"
	params["Disks.0.Type"] = "CLOUD_SSD"
	params["Disks.0.IsBoot"] = "True"
	params["Disks.0.Size"] = int(20)
	params["LoginMode"] = "Password"
	params["Name"] = "ChaosEris_hongwei.liu_20200729_test"
	params["Tag"] = "ChaosEris_hongwei.liu_20200729_test"
	params["SetId"] = uint32(9)
	params["HostIp"] = "10.189.152.130"
	params["MachineType"] = "O"

	params["Memory"]=int(4096)
	params["CPU"]=int(2)
	//params["NetworkInterface.0.EIP.Bandwidth"]=int(40)
	//params["NetworkInterface.0.EIP.OperatorName"]="Bgp"
	//params["NetworkInterface.0.EIP.PayMode"]="Bandwidth"
	var rsp createUHostInstanceResponse
    HTTPReq(params, &rsp)


	fmt.Printf("CreateUHostInstance result:%+v", PrettyPrint(rsp))
	uhost_ids = rsp.UHostIds
	ips = rsp.IPs
	return
}

func HTTPReq(params map[string]interface{}, rsp interface{}) {
	Signature := Verfy_AC(params)
	params["Signature"] = Signature
	externalUrl := "https://api.upc-poc.cn"
	fmt.Printf("Post url: %s, params:%v", externalUrl, PrettyPrint(params))
	b, err := json.Marshal(params)
	if err != nil {
		fmt.Printf("json err: %v", err)
		return
	}
	body := []byte(b)
	rspData, rc := SendHttpPostJsonRequest(externalUrl, body)
	if rc != nil {
		fmt.Printf("http post json request fail; params:%v; rc:%v", params, rc)
		return
	}
	if err := json.Unmarshal(rspData, rsp); err != nil {
		fmt.Printf("unmarshal http rsp fail; err:%v", err)
		return
	}

	vRsp := reflect.ValueOf(rsp).Elem()

	fmt.Println(vRsp)

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


func SendHttpPostJsonRequest(url string, req_body []byte) (rsp_data []byte, rc *ResponseCode) {
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


	var uhost_ids []string
	var uhost_ips []string

		uhost_ids, uhost_ips = CreateUHostInstance()
		time.Sleep(time.Second*20)
		fmt.Println(uhost_ids,uhost_ips)
		fmt.Println("create successful!")
		
}




/*
CreateUHostInstance result:{
	"RetCode": 0,
	"Message": "",
	"Action": "CreateUHostInstanceResponse",
	"UHostIds": [
		"uhost-wopo1piz"
	],
	"IPs": [
		"10.51.39.43"
	]
}[uhost-wopo1piz] [10.51.39.43]
create successful!

*/