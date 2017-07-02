/*
* created by Orz on 2017/6/13
 */
package model

import (
	"encoding/json"
	"im/imserver/log"
	. "im/public"
	"strconv"
)

func QueryRooms(pack *Package, session *Session) {
	log.Log("user:" + session.user.Name + " query chatrooms")

	rooms := GetAllRooms()

	cmdInfo := CmdInfo{
		Cmd: pack.CmdInfo.Cmd,
	}
	if b, err := json.Marshal(rooms); err != nil {
		cmdInfo.Err = ERR_FAILED
		cmdInfo.ErrInfo = "query all rooms failed"
	} else {
		cmdInfo.Data = b
	}

	session.SendMsg(&cmdInfo)
}

func EntryRoom(pack *Package, session *Session) {
	cmdInfo := CmdInfo{
		Cmd: pack.CmdInfo.Cmd,
	}

	if ID, err := strconv.Atoi(string(pack.CmdInfo.Data)); err != nil {
		cmdInfo.Err = ERR_FAILED
		cmdInfo.ErrInfo = "param ID is invalid"
		session.SendMsg(&cmdInfo)
	} else {
		session.chatRoom = GetChatRoom(ID)
		if session.chatRoom == nil {
			cmdInfo.ErrInfo = "the room:" + strconv.Itoa(ID) + " is not exist"
			session.SendMsg(&cmdInfo)
			return
		}
		session.chatRoom.AddNewSession(ID, session)

		cmdInfo.Err = ERR_SUCCESS
		session.SendMsg(&cmdInfo)
		broadcastMsg := "user:" + session.user.Name + " entry room:" + strconv.Itoa(ID)
		log.Log(broadcastMsg)
		broadcastCmd := CmdInfo{
			Cmd:  CHAT,
			Data: []byte(broadcastMsg),
		}
		Broadcast(&broadcastCmd, session)
	}
}
