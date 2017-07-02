// imclient project imclient.go
package main

import (
	"fmt"
	"im/imclient/log"
	"im/imclient/model"
	"im/imclient/network"
	"im/imclient/util"
	"im/imclient/view"
	. "im/public"
)

func main() {
	var sendCh = make(chan *CmdInfo)
	var stopCh = make(chan int)

	network.Connect(sendCh, stopCh)

	fmt.Println(GetCurrTime(), ": connect successfully!")

	model.Login(sendCh)
	if !model.WaitResp() {
		log.Log("query server info timeout")
		return
	}
	view.WaitForInput(sendCh, stopCh)

	util.WaitGroup.Wait()

}
