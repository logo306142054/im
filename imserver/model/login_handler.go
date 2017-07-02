/*
* created by Orz on 2017/6/11
 */
package model

import (
	"encoding/json"
	"im/imserver/log"
	. "im/public"
)

func Login(pack *Package, session *Session) {
	var user User
	if err := json.Unmarshal(pack.CmdInfo.Data, &user); err != nil {
		log.Log("can't parse user info from data")
		return
	}

	user.ID = GetLoginedUsers().AddNewUser(user)
	cmdInfo := CmdInfo{
		Cmd:  LOGIN,
		Data: []byte(user.ID),
	}
	if user.ID != "" {
		session.user = &user
		cmdInfo.Err = ERR_SUCCESS
		log.Log("new user!, ID=" + user.ID + " name=" + user.Name + " " + user.Gender.String())
	} else {
		cmdInfo.Err = ERR_LOGIN_FAILED
		log.Log("login failure!, user=" + user.Name + " " + user.Gender.String())
	}

	session.SendMsg(&cmdInfo)
}

func Logout(pack *Package, session *Session) {
	if session.chatRoom != nil {
		session.chatRoom.DelSession(session.sid)
	}
	if session.user != nil {
		GetLoginedUsers().DelUser(session.user.ID)
	}
}
