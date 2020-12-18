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
	_"time"
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)

}


type copyImageResponse struct {
	ResponseCode
	UHostIds []string
	IPs      []string
}
type ResponseCode struct {
	RetCode        int    `json:"RetCode"`
	TargetImageId  string `json:"TargetImageId"`
	request_uuid   string   `json:"request_uuid"`
	Action         string `json:"Action"` // 只有返回给api调用者时才会填充
}
//创建UHost实例
func CopyCustomImage() (TargetImageId string, request_uuid string,RetCode   int) {
	params := make(map[string]interface{}, 0)
	params["Action"] = "CopyCustomImage"
	params["request_uuid"] = "hongwei_request_20201116"
    params["Region"] = "cn-sh2"
	params["Zone"] = "cn-sh2-01"
	params["ProjectId"] = "org-0sqlid"
	params["SourceImageId"] = "uimage-v512moxl"
	params["TargetProjectId"] = "org-0sqlid"
    params["TargetRegion"] = "cn-bj2"
	params["TargetImageName"] = "clone_image_test_20201116"
	//params["TargetImageDescription"] = "克隆镜像"
	params["PublicKey"] = "OnBypPwlSGnW0NgV6vnghx_lRRgDCruNFL6AywF-"
	//params["PrivateKey"] = "Erd573ujHbfHOUsZdX9mg83qlTxeH9gg4n8nR7C7xn6X5B1kRQG_nN2LSaSGuIyp"
	var rsp copyImageResponse
    HTTPReq(params, &rsp)


	fmt.Printf("CopyCustomImage result:%+v", PrettyPrint(rsp))
	TargetImageId = rsp.TargetImageId
	request_uuid = rsp.request_uuid
	RetCode = rsp.RetCode
	return
}

func HTTPReq(params map[string]interface{}, rsp interface{}) {
	Signature := Verfy_AC(params)
	params["Signature"] = Signature
	externalUrl := "http://api.ucloud.cn"
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
	ParamsData += "Erd573ujHbfHOUsZdX9mg83qlTxeH9gg4n8nR7C7xn6X5B1kRQG_nN2LSaSGuIyp"
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


	func Verfy(params map[string]interface{}) (Signature string) {
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
		ParamsData += "46f09bb9fab4f12dfc160dae12273d5332b5debe"
		h := sha1.New()
		h.Write([]byte(ParamsData))
		bs := h.Sum(nil) // bs 是 16 进制的hash值
		Signature = fmt.Sprintf("%x", bs)
		return Signature
	}

func main(){

	TargetImageId,request_uuid,RetCode := CopyCustomImage()
		fmt.Println(TargetImageId,request_uuid,RetCode)
		fmt.Println("create successful!")
	
	
	/*
	params := make(map[string]interface{}, 0)
	params["PublicKey"] = "ucloudsomeone@example.com1296235120854146120"
	//params["PrivateKey"] = "46f09bb9fab4f12dfc160dae12273d5332b5debe"
	params["Action"] =  "DescribeUHostInstance"
    params["Region"] =  "cn-bj2"
    params["Limit"]   =  10
	fmt.Println(Verfy(params))
	*/
}
