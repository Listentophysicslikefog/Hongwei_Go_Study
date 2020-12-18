/*package main
import (
	//"encoding/json"
	"fmt"
)
func main(){
	//json.MarshalIndent()
}

*/


package main
import (

	"strings"
	"fmt"
)

const (
	ONE_IPTABLES_RULES = int(1)
	IPTABLES_PARAMS_NUM = int(8)
)

type IptablesRules struct {

	PackageLossType     string
	ExecCmdIp           string
	ExecCmdPort         string    
	PeerIp              string
	PeerPort            string
	PacketLossRate      string
	RunningPackageLoss  bool  //是否运行丢包规则
	RunningTime         int    //执行丢包的时间长度，就是停止丢包的周期
	StartPackageLoss    int   //开始执行丢包命令的周期，最大公约数
	FuncTime            string
	RunTime             string    

}

func SplitIptablesRules(allRules string ,rulesNum int)(resolveRes map[string]IptablesRules){
	var splitRes []string
	if rulesNum == ONE_IPTABLES_RULES{
	splitRes = append(splitRes,allRules)
}else {
	splitRes = strings.Split(allRules,",end")
}

resolveRes = make(map[string]IptablesRules,rulesNum +1 )
for _,v:=range splitRes {
oneIptablesRules := strings.Split(v,",")
if len(oneIptablesRules) < IPTABLES_PARAMS_NUM {
	fmt.Println("IpTables Rules Params Error")
}

//var tempRules IptablesRules
key := oneIptablesRules[0] + "_" + oneIptablesRules[1]
fmt.Println(key)
resolveRes[key] = IptablesRules {
	PackageLossType: oneIptablesRules[0],
	ExecCmdIp      : oneIptablesRules[1],
	ExecCmdPort    : oneIptablesRules[2], 
	PacketLossRate : oneIptablesRules[3],
	FuncTime       : oneIptablesRules[4],
	RunTime        : oneIptablesRules[5],
	PeerPort       : oneIptablesRules[6],
	PeerIp         : oneIptablesRules[7],
}

//fmt.Println(resolveRes)
}
return
}

func main(){
	var test ="INPUT1,ExecCmdIp,ExecCmdPort,PackageLossRate,FuncTime,RunTime,PeerPort,PeerIp,endINPUT2,ExecCmdIp,ExecCmdPort,PackageLossRate,FuncTime,RunTime,PeerPort,PeerIp,endINPUT3,ExecCmdIp,ExecCmdPort,PackageLossRate,FuncTime,RunTime,PeerPort,PeerIp,endINPUT4,ExecCmdIp,ExecCmdPort,PackageLossRate,FuncTime,RunTime,PeerPort,PeerIp,endOUTPUT5,ExecCmdIp,ExecCmdPort,PackageLossRate,FuncTime,RunTime,PeerPort,PeerIp"

	re := SplitIptablesRules(test,5)
for k,v:=range re {
	fmt.Printf("key : %v , value : %v\n", k,v )
}

fmt.Println(strings.Split("192.168.2.9_172.3.4.5","_"))
}