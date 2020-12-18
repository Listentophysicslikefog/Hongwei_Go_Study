package main
import (
	//	"fmt"
	"encoding/json"
	"fmt"
	"strings"
	"strconv"
)
//https://golang.org/pkg/encoding/json/#example_Decoder_Decode_stream
type PcMetaInfo struct {
	Is_Used       int
	Has_Detection int
	Seq_No        int
	Pc_Id         int
	Offset        uint
	Pg_Id         int
	Lc_Id         int
	Pc_No         int
	Lc_Random_Id  int
	Allocate_Time uint64
	Md5           string
}
const pc_info_json=`
[
	{"is_used":1,"has_detection":1,"seq_no":2574,"pc_id":164672,"offset":691456503808,"pg_id":41,"lc_id":172,"pc_no":8822,"lc_random_id":217827144,"allocate_time":1589538157,"md5":"e8a8c6233a96e1d56eaa21a4cc62737a"},
	{"is_used":1,"has_detection":1,"seq_no":26396,"pc_id":164673,"offset":691460702208,"pg_id":39,"lc_id":63,"pc_no":24057,"lc_random_id":682186828,"allocate_time":1589538434,"md5":"0837e5889c2e7e9089bf6a314aac563a"},
	{"is_used":1,"has_detection":1,"seq_no":50218,"pc_id":164674,"offset":691464900608,"pg_id":47,"lc_id":63,"pc_no":9379,"lc_random_id":682186828,"allocate_time":1589538696,"md5":"9e2c06532674c0f685438547f55bca06"},
	{"is_used":1,"has_detection":1,"seq_no":2575,"pc_id":164736,"offset":691725201408,"pg_id":43,"lc_id":175,"pc_no":330,"lc_random_id":1921771860,"allocate_time":1589538157,"md5":"944e49a4fe4b91045106a331c7debd78"}
]
	`
func main(){
	
	dec := json.NewDecoder(strings.NewReader(pc_info_json))
	_, err := dec.Token()
	if err != nil {
		panic(err)
	}

	var pg_pc_dict map[int]map[string]string
	pg_pc_dict = make(map[int]map[string]string)
	for dec.More() {
		var pc_meta_info PcMetaInfo
		err := dec.Decode(&pc_meta_info)
		if err != nil {
			fmt.Println(err)
			//panic(err)
		}
		_, ok := pg_pc_dict[pc_meta_info.Pg_Id]
		if !ok { //不存在
			pg_pc_dict[pc_meta_info.Pg_Id] = make(map[string]string)
		}
		lcid_pcno_lcrandomid := strconv.Itoa(pc_meta_info.Lc_Id) + "_" + strconv.Itoa(pc_meta_info.Pc_No) + "_" + strconv.Itoa(pc_meta_info.Lc_Random_Id)
		pg_pc_dict[pc_meta_info.Pg_Id][lcid_pcno_lcrandomid] = pc_meta_info.Md5
		fmt.Printf("lcid_pcno_lcrandomid: %s, md5: %s\n", lcid_pcno_lcrandomid, pc_meta_info.Md5)
		fmt.Println(pc_meta_info)
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		panic(err)
	}

}