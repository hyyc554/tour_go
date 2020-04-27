package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
	"tour_go/biu-master"
)

const UDP_HEAD_LENGTH = 17

type DataContainer struct {
	NetHead   uint8
	PkgLength uint8
	SeqNum    uint8
	GateRssi  int8
	GateId    uint16
	//基站数据
	MsgLength    uint8
	Command      uint8
	AnchorSeq    uint8
	BatLevel     uint8
	TagId        uint16
	Timestamp    uint16
	TimeInterval uint8
	DataCount    uint8
	MoveFlag     uint8
	//混沌数据
	RtagRssi map[int]int
	//混沌数据
	AnchorCrc16 uint16
	//基站数据
	DataCrc16 uint16
	DataEnd   uint8
}

func main() {

	data := DataContainer{}
	b := []byte{0x59, 0x30, 0x21, 0xbf, 0x00, 0x00, 0x2a, 0x88, 0xf6, 0x00, 0x0e, 0x00, 0xf6, 0x01, 0x20, 0x01, 0x02, 0x00, 0x04, 0xae, 0x00, 0x69, 0x5c, 0x01, 0xc6, 0xc6, 0x65, 0xe5, 0x90, 0x07, 0x0b, 0x20, 0x0d, 0x36, 0x40, 0x1c, 0xed, 0x00, 0x04, 0xda, 0x00, 0x6e, 0xb4, 0x00, 0x8d, 0x68, 0x6a, 0x3f, 0x2f, 0x66, 0x47}
	total_len := len(b)
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.LittleEndian, &data.NetHead)
	binary.Read(buf, binary.LittleEndian, &data.PkgLength)
	binary.Read(buf, binary.LittleEndian, &data.SeqNum)
	binary.Read(buf, binary.LittleEndian, &data.GateRssi)
	binary.Read(buf, binary.LittleEndian, &data.GateId)
	binary.Read(buf, binary.LittleEndian, &data.MsgLength)
	binary.Read(buf, binary.LittleEndian, &data.Command)
	binary.Read(buf, binary.LittleEndian, &data.AnchorSeq)
	binary.Read(buf, binary.LittleEndian, &data.BatLevel)
	binary.Read(buf, binary.LittleEndian, &data.TagId)
	binary.Read(buf, binary.LittleEndian, &data.Timestamp)
	binary.Read(buf, binary.LittleEndian, &data.TimeInterval)
	binary.Read(buf, binary.LittleEndian, &data.DataCount)
	binary.Read(buf, binary.LittleEndian, &data.MoveFlag)

	bit_data := fmtBits(b)
	fmt.Printf("<%s>\n", fmtBits(b))
	fmt.Println(len(bit_data))
	start := UDP_HEAD_LENGTH * 8
	iobt_len := total_len*8 - 24 - 16
	rssi_res := map[uint16]map[int]int{}
	for i := start; i+7+16 < iobt_len; i = start {
		end := start + 16
		beacon_bit := bit_data[start:end]
		var beacon_id uint16
		biu.ReadBinaryString(beacon_bit, &beacon_id)
		beaconRes := map[int]int{}
		pos_head := start + 16
		pos_end := pos_head + int(data.DataCount)
		data_pos := bit_data[pos_head:pos_end]
		for i, k := range data_pos {
			if k == 49 {
				var num uint8
				biu.ReadBinaryString(bit_data[pos_end:(pos_end+6)], &num)
				//fmt.Println(bit_data[pos_end:(pos_end+6)], num)
				beaconRes[i] = int(num)*(-1) - 40
				pos_end += 6
			}
		}
		rssi_res[beacon_id] = beaconRes
		start = pos_end

	}
	fmt.Println(rssi_res)

	//fmt.Println(BeaconId) //259
	//k := string(b)
	//fmt.Printf("<%s>\n",k)
	//println(hex.Dump(b))

}

//byte转二进制字符串
func fmtBits(data []byte) string {
	var build strings.Builder
	for _, b := range data {
		s := fmt.Sprintf("%08b", b)
		build.WriteString(s)
	}
	return build.String()
}

//type IobContainer struct {

//}
//
//func readbodata(IobData string,dataContainer *DataContainer) *DataContainer {
//	var BeaconId uint16
//	biu.ReadBinaryString(IobData[:16], &BeaconId)
//	return
//}
