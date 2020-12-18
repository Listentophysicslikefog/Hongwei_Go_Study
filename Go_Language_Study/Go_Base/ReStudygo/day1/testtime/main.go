package main
import (
	"fmt"
	"time"
)

type RepairChunkInfo struct {
  Repair_Detail                []OneChunkRepairInfo
}

type OneChunkRepairInfo struct{
	Repair_Begin_Time             string
	Repair_End_time               string
	Three_Copies_Check_Begin_Time string
	Three_Copies_Check_End_Time   string
	Three_Copies_Check_Result     string
	Bad_Chunk_Id                  uint32
	New_Chunk_Id                  uint32
  }

func main(){

	var repair_chunk_info RepairChunkInfo
	var one_repair_chunk OneChunkRepairInfo 
    one_repair_chunk.Bad_Chunk_Id = uint32(9)
    one_repair_chunk.New_Chunk_Id = uint32(99)
    timestamp:= time.Now().Unix()
    tm := time.Unix(timestamp, 0)
    one_repair_chunk.Repair_Begin_Time = tm.Format("2006-01-02 15:04:05 PM")
    repair_chunk_info.Repair_Detail = append(repair_chunk_info.Repair_Detail, one_repair_chunk)
time.Sleep(time.Second*9)
fmt.Println(len(repair_chunk_info.Repair_Detail))
	for _, one_repair_chunk_detail := range repair_chunk_info.Repair_Detail{
		
		  timestamp:= time.Now().Unix()
		  tm:= time.Unix(timestamp, 0)
		  one_repair_chunk_detail.Three_Copies_Check_End_Time = tm.Format("2006-01-02 15:04:05 PM")
		  one_repair_chunk_detail.Three_Copies_Check_Result = "COMMON_TEST_SUCCEED"

		  repair_chunk_info.Repair_Detail[0].Three_Copies_Check_End_Time = tm.Format("2006-01-02 15:04:05 PM")
		  repair_chunk_info.Repair_Detail[0].Three_Copies_Check_Result = "COMMON_TEST_SUCCEED"
		fmt.Println(repair_chunk_info)
}

         fmt.Println(repair_chunk_info)
}