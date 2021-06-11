package main

import (
	"fmt"
	. "github.com/rfyiamcool/go-shell"
	"time"
)

func main()  {
	cmd := NewCommand("/home/yance/Documents/tour_go/go_sh_demo/go-shell/sleep.sh", WithTimeout(10), WithShellMode())
	cmd.Start()
	time.Sleep(3*time.Second)
	status := cmd.Status
	fmt.Println(status.PID)
	cmd.Wait()


}
