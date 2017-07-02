/*
* created by Orz on 2017/6/13
 */
package model

import (
	"im/imserver/log"
	. "im/public"
)

func Say(ct *ChatRoom, msg string) {

}

func Broadcast(cmdInfo *CmdInfo, session *Session) {
	if session.chatRoom != nil {
		for _, s := range session.chatRoom.sessions {
			s.SendMsg(cmdInfo)
		}
	}
}

func Chat(pack *Package, session *Session) {
	msg := pack.Time + " " + session.user.Name + " say:" + string(pack.CmdInfo.Data)
	pack.CmdInfo.Data = []byte(msg)
	log.Log(msg)
	Broadcast(&pack.CmdInfo, session)

}
