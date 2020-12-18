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


func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)

}

type attachUDiskResponse struct {
	ResponseCode
	UDiskId string
	UHostId string
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

	
	
	func AttachUDisk() (uhost_id string, udisk_id string, rc *ResponseCode) {
		params := make(map[string]interface{}, 0)
		params["Action"] = "AttachUDisk"
		params["request_uuid"] = "hongwei_request"
		params["Region"] = "cn-gd2"
		params["Zone"] = "cn-gd2-01"
		params["ProjectId"] = "org-nf3wlo"
		params["PublicKey"] = "O9aQkG61gppmkiPhiZyuljs+2yPwcmSXRMX9wI7yZujZ/qlOPKpPew=="
		params["UHostId"] = "uhost-m1z1eymb"  //uhost-m1z1eymb                                   //uhost-vin5bsiu   uhost-wopo1piz  uhost-nogjrntq
		params["UDiskId"] = "bsr-5a3jv2j0"   //   bsr-5a3jv2j0                                                            //bsr-ysahqv5c  bsr-yp0brqvn  bsr-q2ogp344   bsr-kr2cibpu   bsr-hluxvkuw  bsr-u43hi0jv  bsr-2ov24joo
		params["MultiAttach"] = "No"
	
		var rsp attachUDiskResponse
		HTTPReq(params, &rsp)
	
	
		fmt.Printf("attach udisk result:%+v", PrettyPrint(rsp))
		udisk_id = rsp.UDiskId
		uhost_id = rsp.UHostId
		return
	}

	func AttachUDisks(attachudiskid string) (uhost_id string, udisk_id string, rc *ResponseCode) {
		params := make(map[string]interface{}, 0)
		params["Action"] = "AttachUDisk"
		params["request_uuid"] = "hongwei_request"
		params["Region"] = "cn-gd2"
		params["Zone"] = "cn-gd2-01"
		params["ProjectId"] = "org-nf3wlo"
		params["PublicKey"] = "O9aQkG61gppmkiPhiZyuljs+2yPwcmSXRMX9wI7yZujZ/qlOPKpPew=="
		params["UHostId"] = "uhost-cc5lh4ym"  //uhost-cc5lh4ym  uhost-o0ng5vg2 / uhost-xeeshf5l   uhost-m1z1eymb   uhost-zzvlmd42    uhost-lwhrc4jp     uhost-s03f1ull                       //uhost-vin5bsiu   uhost-wopo1piz  uhost-nogjrntq
		params["UDiskId"] = attachudiskid   //   bsr-5a3jv2j0                                                            //bsr-ysahqv5c  bsr-yp0brqvn  bsr-q2ogp344   bsr-kr2cibpu   bsr-hluxvkuw  bsr-u43hi0jv  bsr-2ov24joo
		params["MultiAttach"] = "No"
	
		var rsp attachUDiskResponse
		HTTPReq(params, &rsp)
	
	
		fmt.Printf("attach udisk result:%+v", PrettyPrint(rsp))
		udisk_id = rsp.UDiskId
		uhost_id = rsp.UHostId
		return
	}

	func HTTPReq(params map[string]interface{}, rsp interface{}) {
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
		var uhost_ids string
		var udisk_ids string

		//uhost_ids,udisk_ids,rc=AttachUDisk()
	   var udisk []string
	  // udisk=[]string{"bsr-gdyxek1r","bsr-jomrv1yv","bsr-34tepohj","bsr-3nkizjrv","bsr-c2bd4w3e","bsr-j4ono1dz","bsr-1f2z04lx","bsr-bgw0kve5","bsr-3bg2veyq","bsr-gsglgrcf","bsr-arqinih4","bsr-sim3ftip","bsr-fvnlha3l","bsr-li3psk4z","bsr-kyjeasw5","bsr-zts0y2cv","bsr-4xu3bnw2","bsr-3sebbr5l"}
	//	udisk=[]string{"bsr-aqc14nkh","bsr-3yyijpgt","bsr-0dirgjke","bsr-kktvyj13","bsr-wthym2ef","bsr-rdcqbonn","bsr-0tcp2pn2","bsr-vko3lzav","bsr-yiycw1bl","bsr-kiktdlup","bsr-uuwgpfqu","bsr-n5qgede0","bsr-wvfxveku","bsr-knkrnorh","bsr-ikcluc2k","bsr-55z3sqnt","bsr-i21du1gs","bsr-s1isge32"}
	 udisk=[]string{"bsr-yobarlay","bsr-1pempx3l","bsr-w1szvfqg","bsr-md3ix4jp","bsr-uf2wo3u3","bsr-iqk55aei","bsr-w3dz1l1w","bsr-phttn4ci","bsr-huyecwnj","bsr-zg41hrsr","bsr-efo33e0x","bsr-4sux4z04","bsr-immfen1l","bsr-14zedwdv","bsr-eoivw4vb","bsr-yhwy4i30","bsr-rh1gvxhp","bsr-ybqwmyx0"}
	  for k,v:=range udisk{
			uhost_ids,udisk_ids,rc=AttachUDisks(v)
			fmt.Println(k)
			time.Sleep(time.Second*20)
			fmt.Println(udisk_ids,"   ", uhost_ids ,rc)
		}
		
		
}
