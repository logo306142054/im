package model

import (
	"im/imclient/log"
	. "im/public"
)

func Login(sendCh chan *CmdInfo) {
	Host.Name = "chx"
	Host.Gender = MALE
	cmdInfo := CmdInfo{
		Cmd:  LOGIN,
		Data: ParseToJsonBytes(Host),
	}

	sendCh <- &cmdInfo
}

func RespLogin(pack *Package, sendCh chan *CmdInfo) {
	if pack.CmdInfo.Err == ERR_SUCCESS {
		Host.ID = string(pack.CmdInfo.Data)
		log.Log("login successfully!")
	} else {
		log.Log("login failure!")
	}

	QueryRooms(sendCh)
}
