package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/imroc/biu"

)

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
	RSSI []byte
	//混沌数据
	AnchorCrc16 uint16
	//基站数据
	DataCrc16 uint16
	DataEnd   uint8
}

func main() {

	data := DataContainer{}
	b := []byte{0x59, 0x30, 0x21, 0xbf, 0x00, 0x00, 0x2a, 0x88, 0xf6, 0x00, 0x0e, 0x00, 0xf6, 0x01, 0x20, 0x01, 0x02, 0x00, 0x04, 0xae, 0x00, 0x69, 0x5c, 0x01, 0xc6, 0xc6, 0x65, 0xe5, 0x90, 0x07, 0x0b, 0x20, 0x0d, 0x36, 0x40, 0x1c, 0xed, 0x00, 0x04, 0xda, 0x00, 0x6e, 0xb4, 0x00, 0x8d, 0x68, 0x6a, 0x3f, 0x2f, 0x66, 0x47}
	data_length := len(b)
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
	fmt.Println(b)
	//if err != nil {
	//	log.Fatalln("binary.Read failed:", err)
	//}


	fmt.Println(s)

	fmt.Println(data)

}
