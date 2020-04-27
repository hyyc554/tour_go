package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	DLLPATH = `C:\Users\YC\Documents\go_project\tour_go\gowithc\lib\hasp_windows_x64_demo.dll`
)

func IntPtr(n int) uintptr {
	return uintptr(n)
}

func StrPtr(s string) uintptr {
	//b := append([]byte(s), 0)
	s16, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		fmt.Println(err)
	}
	//s16:=syscall.StringToUTF16Ptr(s)
	return uintptr(unsafe.Pointer(s16))
}

func WindowVersion2() {
	dll32 := syscall.NewLazyDLL("kernel32.dll")

	g := dll32.NewProc("GetVersion")

	r, _, _ := g.Call()
	major := byte(r)
	minor := uint8(r >> 8)
	build := uint16(r >> 16)
	print("windows version ", major, ".", minor, " (Build ", build, ")\n")
}

func Hello3() {
	var feature uint64
	var HASP_INVALID_HANDLE_VALUE uint64
	var vendorCode string
	//var vendorCode1 string
	//vendorCode = "XBFph/yyR3E6zeNjULioIpwrksI/DvorlpkdXWOyjZmAvILLpqJ3wAQTiX4O45IpJCVADfaYzmcZzcfJ"
	//vendorCode1 = "8ea2sJ9DnZyYVBTy+IrTX1TkdLO9xeNCMCCCxvTFPjai7x5yYiI2E8sMdH3xmKM5yBUtcAqHridGEfDH"
	//vendorCode = vendorCode+vendorCode1
	//vendorCode = "XBFph/yyR3E6zeNjULioIpwrksI/DvorlpkdXWOyjZmAvILLpqJ3wAQTiX4O45IpJCVADfaYzmcZzcfJ8ea2sJ9DnZyYVBTy+IrTX1TkdLO9xeNCMCCCxvTFPjai7x5yYiI2E8sMdH3xmKM5yBUtcAqHridGEfDHZKTIiDP+hxwmKc4wXoH+3XJJr5QF5ib/aZNnU93hJXLAdEZ+zXCUF/CC929L4GFvgoLev6T42lhonz+0bMtThu7Cwo6UVRuFfniIeIJH6nLhu9YBrkgXEWouw7z+s3M5tIvS11iFCHbqrcxnXgFKnaAfWOTKTNhWPXmGpxta6eZEuJgUNiLySs1VwyGOxp0F1j0IiKK7g6QMYbjSCemCwTCWuxBdQucAUojAy+B2BMhBq/dEjvGXCkTmU08qI3AF2LAnBsxFF4gNfSKPapADCD1PQDUDHVD6oqhtj5HkDakczIjBuwQUykmjI0MQzdsy7VYNsfHI2qQYreIcES5RQEqjX89zyeYFw0B8hh1oM0rE8aNO3kCH0EL4FKeqalgNgOxo2I+3eweCjzB9+B9BD3UQHvf8qf25KXQyUSoNPjPIJRaHi/8NomOmznP8Hy8pYcnqMCoC847cBXAadPeCqN27znG2YlM+vSjkmEEGcZCcZA1urdeYRyGZdbnG5YmTnccNkO4Nol3r7U+oz9CTSvJTvT6es9n55VWRk9tyKASmURchjaHtVdRcW5QglmZ6XtzfSjnNVUKSl3Jh/DJTyq7v9cAw5zpTPjZJdC+8kDX8KfEooiTHWB4HIjaUw9xPSS04IZ9ZQoCUp2fW6rEv2Dm7TMDCXUYP8xW7TcSemoKMeqQjhJBdvR/uVgg5MNtBagxESbP6sLMPIBDOf+/ety998pYEZpGVz/UGwOkIupzMc7G+RyjeunP6w9hpQSdWVjfHA2F/9Dw="
	vendorCode = "XBFph/yyR3E6zeNjULioIpwrksI/DvorlpkdXWOyjZmAvILLpqJ3wAQTiX4O45IpJCVADfaYzmcZzcfJ" +
		"8ea2sJ9DnZyYVBTy+IrTX1TkdLO9xeNCMCCCxvTFPjai7x5yYiI2E8sMdH3xmKM5yBUtcAqHridGEfDH" +
		"ZKTIiDP+hxwmKc4wXoH+3XJJr5QF5ib/aZNnU93hJXLAdEZ+zXCUF/CC929L4GFvgoLev6T42lhonz+0" +
		"bMtThu7Cwo6UVRuFfniIeIJH6nLhu9YBrkgXEWouw7z+s3M5tIvS11iFCHbqrcxnXgFKnaAfWOTKTNhW" +
		"PXmGpxta6eZEuJgUNiLySs1VwyGOxp0F1j0IiKK7g6QMYbjSCemCwTCWuxBdQucAUojAy+B2BMhBq/dE" +
		"jvGXCkTmU08qI3AF2LAnBsxFF4gNfSKPapADCD1PQDUDHVD6oqhtj5HkDakczIjBuwQUykmjI0MQzdsy" +
		"7VYNsfHI2qQYreIcES5RQEqjX89zyeYFw0B8hh1oM0rE8aNO3kCH0EL4FKeqalgNgOxo2I+3eweCjzB9" +
		"+B9BD3UQHvf8qf25KXQyUSoNPjPIJRaHi/8NomOmznP8Hy8pYcnqMCoC847cBXAadPeCqN27znG2YlM+" +
		"vSjkmEEGcZCcZA1urdeYRyGZdbnG5YmTnccNkO4Nol3r7U+oz9CTSvJTvT6es9n55VWRk9tyKASmURch" +
		"jaHtVdRcW5QglmZ6XtzfSjnNVUKSl3Jh/DJTyq7v9cAw5zpTPjZJdC+8kDX8KfEooiTHWB4HIjaUw9xP" +
		"SS04IZ9ZQoCUp2fW6rEv2Dm7TMDCXUYP8xW7TcSemoKMeqQjhJBdvR/uVgg5MNtBagxESbP6sLMPIBDO" +
		"f+/ety998pYEZpGVz/UGwOkIupzMc7G+RyjeunP6w9hpQSdWVjfHA2F/9Dw="

	HASP_INVALID_HANDLE_VALUE = 0
	helloDll := syscall.NewLazyDLL(DLLPATH)
	hello := helloDll.NewProc("hasp_login")
	feature = 1
	r, _, _ := hello.Call(uintptr(feature), StrPtr(vendorCode), uintptr(unsafe.Pointer(&HASP_INVALID_HANDLE_VALUE)))
	major := uint32(r)
	//print(major)
	fmt.Println(major)

}
func main() {
	Hello3()
	WindowVersion2()
}
