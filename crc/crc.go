package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/sigurn/crc16"
	"hash/crc32"
)

// 生成md5
func MD5(str string) string {
	c := md5.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

//生成sha1
func SHA1(str string) string {
	c := sha1.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

func CRC32(str string) uint32 {
	checksum := crc32.ChecksumIEEE([]byte(str))
	fmt.Printf("check sum:%X \n", checksum)
	return checksum
}

func Crc16(str string) uint16 {

	table := crc16.MakeTable(crc16.CRC16_MODBUS)
	crc := crc16.Checksum([]byte(str), table)
	fmt.Printf("CRC-16 MAXIM: %X \n", crc)
	return crc
}
func byCrc16(data []byte) uint16 {
	fmt.Println(data)
	table := crc16.MakeTable(crc16.CRC16_MODBUS)
	crc := crc16.Checksum(data, table)
	fmt.Printf("CRC-16 MAXIM: %X \n", crc)
	return crc
}
func main() {

	//crc_msg_body = by_data[6:-5]
	//crc16_res = DataEngine.crc16_test(crc_msg_body)
	b := []byte{0x59, 0x30, 0x21, 0xbf, 0x00, 0x00, 0x2a, 0x88, 0xf6, 0x00, 0x0e, 0x00, 0xf6, 0x01, 0x20, 0x01, 0x02, 0x00, 0x04, 0xae, 0x00, 0x69, 0x5c, 0x01, 0xc6, 0xc6, 0x65, 0xe5, 0x90, 0x07, 0x0b, 0x20, 0x0d, 0x36, 0x40, 0x1c, 0xed, 0x00, 0x04, 0xda, 0x00, 0x6e, 0xb4, 0x00, 0x8d, 0x68, 0x6a, 0x3f, 0x2f, 0x66, 0x47}

	//fmt.Println(CRC32("test"))
	//fmt.Println(MD5("123456"))
	//fmt.Println(SHA1("123456"))

	fmt.Println(byCrc16(b[6 : len(b)-5]))
}
