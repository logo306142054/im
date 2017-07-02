package model

import (
	"encoding/json"
	"im/imclient/log"
	. "im/public"
)

func QueryRooms(sendCh chan *CmdInfo) {
	cmdInfo := CmdInfo{
		Cmd: QUERY_ROOMS,
	}

	sendCh <- &cmdInfo
}

func RespQueryRooms(pack *Package, sendCh chan *CmdInfo) {
	if pack.CmdInfo.Err != ERR_SUCCESS {
		log.Log(pack.CmdInfo.ErrInfo)
		waitCh <- ERR_FAILED
		return
	}

	var tmp = make([]*Room, 0)
	err := json.Unmarshal(pack.CmdInfo.Data, &tmp)
	if err != nil {
		log.Log("parse rooms failed!")
		waitCh <- ERR_FAILED
		return
	}

	Server.UpdateRooms(tmp)
	waitCh <- 0
}

func RespEntryRoom(pack *Package, sendCh chan *CmdInfo) {
	if pack.CmdInfo.Err != ERR_SUCCESS {
		log.Log(pack.CmdInfo.ErrInfo)
		waitCh <- pack.CmdInfo.Err
		return
	}

	waitCh <- pack.CmdInfo.Err
}

func RespBroadcast(pack *Package, sendCh chan *CmdInfo) {
	if pack.CmdInfo.Err != ERR_SUCCESS {
		return
	}

	CurChatRoom.ReceiveMsg(string(pack.CmdInfo.Data))
}
