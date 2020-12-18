package main
import (
	"fmt"
)
type block struct{
	data interface{}
}

func main(){
	var ma=map[string]map[string]interface{}{
		"chaosgo_conf": make(map[string]interface{}),
	}
	
	ma["chaosgo_conf"]["project_id"]={
	"log_dir": "/root/chaos/chaos-hn08-5009/log",
    "log_prefix": "vm_controller",
    "log_suffix": ".log",
    "log_size": 50,
    "log_level": "DEBUG",
	"log_func_file_line": true
	}
	var blo *block
	res:=&blo{
		data:ma,
	}
	fmt.Println(res)

}