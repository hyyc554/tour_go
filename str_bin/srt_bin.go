package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

func strToBinary(s string, base int) []byte {

	var b []byte

	for _, c := range s {
		b = strconv.AppendInt(b, int64(c), base)
	}

	return b
}
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
func main() {
	a := "0101"
	//a := []
	//n := int64(123)
	//
	//a:= strconv.FormatInt(n, 2)
	k := strToBinary(a, 2)
	fmt.Println(k)

	//fmt.Println(IntToBytes(25733))
	//if i, err := strconv.ParseInt("1001", 2, 64); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(i)
	//}
}
