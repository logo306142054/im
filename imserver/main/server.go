package main

import (
	"im/imserver/model"
	"im/imserver/network"
	"im/imserver/util"
)

func main() {

	ct := model.NewChatRoom("sys room", "default room", "system")
	model.SaveChatRoom(ct)

	network.OpenSocket()
	util.WaitGroup.Wait()
}
